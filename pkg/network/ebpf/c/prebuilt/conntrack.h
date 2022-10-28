#ifndef __CONNTRACK_H
#define __CONNTRACK_H

#include <net/netfilter/nf_conntrack.h>
#include <linux/types.h>
#include <linux/sched.h>
#include "bpf_builtins.h"
#include "tracer.h"
#include "conntrack-types.h"
#include "conntrack-maps.h"
#include "ip.h"
#include "ipv6.h"
#include "sock.h"

static __always_inline __u32 systemprobe_pid() {
    __u64 val = 0;
    LOAD_CONSTANT("systemprobe_pid", val);
    return (__u32)val;
}

static __always_inline u32 ct_status(const struct nf_conn *ct) {
    u32 status = 0;
    bpf_probe_read_kernel(&status, sizeof(status), (void *)&ct->status);
    return status;
}

// depending on the kernel version p_net may be a struct net** or possible_net_t*
static __always_inline u32 get_netns(void *p_net) {
    u32 net_ns_inum = 0;
    struct net *ct_net = NULL;
    bpf_probe_read_kernel(&ct_net, sizeof(ct_net), p_net);
    log_debug("prebuilt netns offset: %u\n", offset_ino());
    bpf_probe_read_kernel(&net_ns_inum, sizeof(net_ns_inum), ((char*)ct_net) + offset_ino());
    return net_ns_inum;
}

static __always_inline void print_translation(const conntrack_tuple_t *t) {
    if (t->metadata & CONN_TYPE_TCP) {
        log_debug("TCP\n");
    } else {
        log_debug("UDP\n");
    }

    print_ip(t->saddr_h, t->saddr_l, t->sport, t->metadata);
    print_ip(t->daddr_h, t->daddr_l, t->dport, t->metadata);
}

static __always_inline int nf_conntrack_tuple_to_conntrack_tuple(conntrack_tuple_t *t, const struct nf_conntrack_tuple *ct) {
    bpf_memset(t, 0, sizeof(conntrack_tuple_t));

    switch (ct->dst.protonum) {
    case IPPROTO_TCP:
        t->metadata = CONN_TYPE_TCP;
        t->sport = ct->src.u.tcp.port;
        t->dport = ct->dst.u.tcp.port;
        break;
    case IPPROTO_UDP:
        t->metadata = CONN_TYPE_UDP;
        t->sport = ct->src.u.udp.port;
        t->dport = ct->dst.u.udp.port;
        break;
    default:
        log_debug("ERR(to_conn_tuple): unknown protocol number: %u\n", ct->dst.protonum);
        return 0;
    }

    t->sport = bpf_ntohs(t->sport);
    t->dport = bpf_ntohs(t->dport);
    if (t->sport == 0 || t->dport == 0) {
        log_debug("ERR(to_conn_tuple): src/dst port not set: src: %u, dst: %u\n", t->sport, t->dport);
        return 0;
    }

    if (ct->src.l3num == AF_INET) {
        t->metadata |= CONN_V4;
        t->saddr_l = ct->src.u3.ip;
        t->daddr_l = ct->dst.u3.ip;

        if (!t->saddr_l || !t->daddr_l) {
            log_debug("ERR(to_conn_tuple.v4): src/dst addr not set src:%u, dst:%u\n", t->saddr_l, t->daddr_l);
            return 0;
        }
    } else if (is_ipv6_enabled() && ct->src.l3num == AF_INET6) {
        t->metadata |= CONN_V6;
        read_in6_addr(&t->saddr_h, &t->saddr_l, &ct->src.u3.in6);
        read_in6_addr(&t->daddr_h, &t->daddr_l, &ct->dst.u3.in6);

        if (!(t->saddr_h || t->saddr_l)) {
            log_debug("ERR(to_conn_tuple.v6): src addr not set: src_l: %llu, src_h: %llu\n",
                t->saddr_l, t->saddr_h);
            return 0;
        }
        if (!(t->daddr_h || t->daddr_l)) {
            log_debug("ERR(to_conn_tuple.v6): dst addr not set: dst_l: %llu, dst_h: %llu\n",
                t->daddr_l, t->daddr_h);
            return 0;
        }
    }

    return 1;
}

static __always_inline void increment_telemetry_registers_count() {
    u64 key = 0;
    conntrack_telemetry_t *val = bpf_map_lookup_elem(&conntrack_telemetry, &key);
    if (val == NULL) {
        return;
    }
    __sync_fetch_and_add(&val->registers, 1);
}

static __always_inline int nf_conn_to_conntrack_tuples(struct nf_conn* ct, conntrack_tuple_t* orig, conntrack_tuple_t* reply) {
    struct nf_conntrack_tuple_hash tuplehash[IP_CT_DIR_MAX];
    bpf_memset(tuplehash, 0, sizeof(tuplehash));
    bpf_probe_read_kernel(&tuplehash, sizeof(tuplehash), &ct->tuplehash);

    struct nf_conntrack_tuple orig_tup = tuplehash[IP_CT_DIR_ORIGINAL].tuple;
    struct nf_conntrack_tuple reply_tup = tuplehash[IP_CT_DIR_REPLY].tuple;

    u32 netns = get_netns(&ct->ct_net);

    if (!nf_conntrack_tuple_to_conntrack_tuple(orig, &orig_tup)) {
        return 1;
    }
    orig->netns = netns;

    log_debug("orig\n");
    print_translation(orig);

    if (!nf_conntrack_tuple_to_conntrack_tuple(reply, &reply_tup)) {
        return 1;
    }
    reply->netns = netns;

    log_debug("reply\n");
    print_translation(reply);

    return 0;
}

#endif
