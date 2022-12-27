#ifndef __HTTP2_DEFS_H
#define __HTTP2_DEFS_H

#include <linux/types.h>

// A limit of max frames we will upload from a single connection to the user mode.
// NOTE: we may need to revisit this const if we need to capture more connections.
#define HTTP2_MAX_FRAMES 5

// A limit of max frame size in order to be able to load a max size and pass the varifier.
// NOTE: we may need to change the max size.
#define HTTP2_MAX_FRAME_LEN 10

// A limit of max frame size in order to be able to load a max size and pass the varifier.
// NOTE: we may need to change the max size.
#define HTTP2_MAX_PATH_LEN 32

typedef enum {
    kAuthority = 1,
    kMethod = 2,
    kPath = 4,
    kScheme = 6,
    kStatus = 9,
} __attribute__ ((packed)) header_key;

typedef enum {
    kGET = 2,
    kPOST = 3,
    kEmptyPath = 4,
    kIndexPath = 5,
    kHTTP = 6,
    kHTTPS = 7,
    k200 = 8,
    k204 = 9,
    k206 = 10,
    k304 = 11,
    k400 = 12,
    k404 = 13,
    k500 = 14,
} __attribute__ ((packed)) header_value;

typedef struct {
    header_key name;
    header_value value;
} static_table_value;

typedef struct {
    char path_buffer[32] __attribute__ ((aligned (8)));
} __attribute__ ((packed)) dynamic_string_value;

typedef struct {
    __u64 index;
    dynamic_string_value value;
} dynamic_table_value;

#define MAX_STATIC_TABLE_INDEX 64

#endif
