//go:build windows && npm
// +build windows,npm

// Code generated - DO NOT EDIT.
package model

import (
	"github.com/DataDog/datadog-agent/pkg/security/secl/compiler/eval"
	"net"
	"reflect"
)

func (m *Model) GetIterator(field eval.Field) (eval.Iterator, error) {
	switch field {
	case "process.ancestors":
		return &ProcessAncestorsIterator{}, nil
	case "signal.target.ancestors":
		return &ProcessAncestorsIterator{}, nil
	}
	return nil, &eval.ErrIteratorNotSupported{Field: field}
}
func (m *Model) GetEventTypes() []eval.EventType {
	return []eval.EventType{
		eval.EventType("bind"),
		eval.EventType("capset"),
		eval.EventType("chmod"),
		eval.EventType("chown"),
		eval.EventType("dns"),
		eval.EventType("exec"),
		eval.EventType("exit"),
		eval.EventType("link"),
		eval.EventType("mkdir"),
		eval.EventType("mount"),
		eval.EventType("open"),
		eval.EventType("removexattr"),
		eval.EventType("rename"),
		eval.EventType("rmdir"),
		eval.EventType("setgid"),
		eval.EventType("setuid"),
		eval.EventType("setxattr"),
		eval.EventType("signal"),
		eval.EventType("splice"),
		eval.EventType("unlink"),
		eval.EventType("utimes"),
	}
}
func (m *Model) GetEvaluator(field eval.Field, regID eval.RegisterID) (eval.Evaluator, error) {
	switch field {
	case "async":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.Async
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "bind.addr.family":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Bind.AddrFamily)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "bind.addr.ip":
		return &eval.CIDREvaluator{
			EvalFnc: func(ctx *eval.Context) net.IPNet {
				ev := ctx.Event.(*Event)
				return ev.Bind.Addr.IPNet
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "bind.addr.port":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Bind.Addr.Port)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "bind.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Bind.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "capset.cap_effective":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Capset.CapEffective)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "capset.cap_permitted":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Capset.CapPermitted)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.destination.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.destination.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Chmod.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Chmod.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Chmod.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chmod.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chmod.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Chmod.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Chmod.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Chmod.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chmod.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Chmod.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chmod.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chmod.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.destination.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.destination.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveChownGID(ev, &ev.Chown)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.destination.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.destination.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveChownUID(ev, &ev.Chown)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Chown.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Chown.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Chown.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chown.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chown.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Chown.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Chown.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Chown.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "chown.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Chown.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "chown.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Chown.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "container.id":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveContainerID(ev, &ev.ContainerContext)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "container.tags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveContainerTags(ev, &ev.ContainerContext)
			},
			Field:  field,
			Weight: 9999 * eval.HandlerWeight,
		}, nil
	case "dns.id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.DNS.ID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "dns.question.class":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.DNS.Class)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "dns.question.count":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.DNS.Count)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "dns.question.length":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.DNS.Size)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "dns.question.name":
		return &eval.IntEvaluator{
			OpOverrides: eval.DNSNameCmp,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.DNS.Name)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "dns.question.type":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.DNS.Type)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.args":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgs(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "exec.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.args_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "exec.argv0":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "exec.cap_effective":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.CapEffective)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.cap_permitted":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.CapPermitted)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.comm":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.Comm
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.container.id":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.ContainerID
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.cookie":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Cookie)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.created_at":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.Exec.Process))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.egid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.EGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.egroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.Credentials.EGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.envs_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.Exec.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.euid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.EUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.euser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.Credentials.EUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exec.Process.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exec.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exec.Process.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exec.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exec.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exec.Process.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exec.Process.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exec.Process.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exec.Process.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exec.Process.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exec.Process.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exec.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.fsgid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.FSGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.fsgroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.Credentials.FSGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.fsuid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.FSUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.fsuser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.Credentials.FSUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.Credentials.Group
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exec.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.interpreter.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.interpreter.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exec.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exec.is_kworker":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.PIDContext.IsKworker
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.is_thread":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.IsThread
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.pid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.PIDContext.Pid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.ppid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.PPid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.tid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.PIDContext.Tid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.tty_name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.TTYName
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exec.Process.Credentials.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exec.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exec.Process.Credentials.User
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.args":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgs(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "exit.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.args_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "exit.argv0":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "exit.cap_effective":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.CapEffective)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.cap_permitted":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.CapPermitted)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.cause":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Cause)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.code":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Code)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.comm":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.Comm
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.container.id":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.ContainerID
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.cookie":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Cookie)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.created_at":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.Exit.Process))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.egid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.EGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.egroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.Credentials.EGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.envs_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.Exit.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.euid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.EUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.euser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.Credentials.EUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exit.Process.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exit.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exit.Process.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exit.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exit.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exit.Process.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exit.Process.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exit.Process.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exit.Process.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exit.Process.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Exit.Process.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exit.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.fsgid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.FSGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.fsgroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.Credentials.FSGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.fsuid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.FSUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.fsuser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.Credentials.FSUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.Credentials.Group
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exit.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.interpreter.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.interpreter.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Exit.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "exit.is_kworker":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.PIDContext.IsKworker
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.is_thread":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.IsThread
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.pid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.PIDContext.Pid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.ppid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.PPid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.tid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.PIDContext.Tid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.tty_name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.TTYName
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Exit.Process.Credentials.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "exit.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Exit.Process.Credentials.User
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Source.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Target.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Link.Target)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Target.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Link.Target.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Link.Target.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Target.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Target.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Target.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Target.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Target)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Target))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Target)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Target))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Link.Target.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.destination.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Target.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.destination.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Link.Target.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Link.Source)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Source.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Link.Source.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Link.Source.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Source.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Source.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Source.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Source.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Source)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Source))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Source)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Source))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Link.Source.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.Source.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "link.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Link.Source.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "link.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Link.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.destination.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.destination.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Mkdir.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Mkdir.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Mkdir.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Mkdir.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Mkdir.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Mkdir.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Mkdir.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Mkdir.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mkdir.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Mkdir.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mkdir.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mkdir.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mount.fs_type":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Mount.Mount.FSType
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mount.mountpoint.path":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveMountPointPath(ev, &ev.Mount)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "mount.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Mount.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "mount.source.path":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveMountSourcePath(ev, &ev.Mount)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "network.destination.ip":
		return &eval.CIDREvaluator{
			EvalFnc: func(ctx *eval.Context) net.IPNet {
				ev := ctx.Event.(*Event)
				return ev.NetworkContext.Destination.IPNet
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "network.destination.port":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.NetworkContext.Destination.Port)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "network.device.ifindex":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.NetworkContext.Device.IfIndex)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "network.device.ifname":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveNetworkDeviceIfName(ev, &ev.NetworkContext.Device)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "network.l3_protocol":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.NetworkContext.L3Protocol)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "network.l4_protocol":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.NetworkContext.L4Protocol)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "network.size":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.NetworkContext.Size)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "network.source.ip":
		return &eval.CIDREvaluator{
			EvalFnc: func(ctx *eval.Context) net.IPNet {
				ev := ctx.Event.(*Event)
				return ev.NetworkContext.Source.IPNet
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "network.source.port":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.NetworkContext.Source.Port)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.destination.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Open.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Open.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Open.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Open.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Open.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Open.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Open.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Open.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Open.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "open.flags":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.Flags)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "open.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Open.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.ancestors.args":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgs(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: 100 * eval.IteratorWeight,
		}, nil
	case "process.ancestors.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgsFlags(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgsOptions(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.args_truncated":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgv(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: 100 * eval.IteratorWeight,
		}, nil
	case "process.ancestors.argv0":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgv0(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: 100 * eval.IteratorWeight,
		}, nil
	case "process.ancestors.cap_effective":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.CapEffective)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.cap_permitted":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.CapPermitted)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.comm":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Comm
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.container.id":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.ContainerID
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.cookie":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Cookie)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.created_at":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &element.ProcessContext.Process))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.egid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.EGID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.egroup":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.EGroup
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessEnvp(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessEnvs(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.envs_truncated":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.euid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.EUID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.euser":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.EUser
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.change_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.CTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.filesystem":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.gid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.GID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.group":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.in_upper_layer":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, false)
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.inode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.Inode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.mode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.Mode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.modification_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.MTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.mount_id":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.MountID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.name":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.name.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.path":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.path.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.rights":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.FileEvent.FileFields))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.uid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.UID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.file.user":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.fsgid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.FSGID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.fsgroup":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.FSGroup
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.fsuid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.FSUID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.fsuser":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.FSUser
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.gid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.GID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.group":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.Group
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.change_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.filesystem":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.gid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.group":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.in_upper_layer":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, false)
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.inode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.mode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.modification_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.mount_id":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.name":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.name.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.path":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.path.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.rights":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.uid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.interpreter.file.user":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.is_kworker":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.PIDContext.IsKworker
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.is_thread":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.IsThread
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.pid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.PIDContext.Pid)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.ppid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.PPid)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.tid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.PIDContext.Tid)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.tty_name":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.TTYName
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.uid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.UID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.ancestors.user":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.User
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "process.args":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgs(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "process.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsFlags(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsOptions(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.args_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "process.argv0":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv0(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "process.cap_effective":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.CapEffective)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.cap_permitted":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.CapPermitted)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.comm":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.Comm
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.container.id":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.ContainerID
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.cookie":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Cookie)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.created_at":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &ev.ProcessContext.Process))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.egid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.EGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.egroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.Credentials.EGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvp(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvs(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.envs_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &ev.ProcessContext.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.euid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.EUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.euser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.Credentials.EUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Process.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Process.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Process.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Process.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Process.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Process.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Process.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Process.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.fsgid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.FSGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.fsgroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.Credentials.FSGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.fsuid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.FSUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.fsuser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.Credentials.FSUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.Credentials.Group
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.interpreter.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.interpreter.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.is_kworker":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.PIDContext.IsKworker
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.is_thread":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.IsThread
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.args":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.FieldHandlers.ResolveProcessArgs(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "process.parent.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.args_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return false
				}
				return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessArgv(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "process.parent.argv0":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "process.parent.cap_effective":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.CapEffective)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.cap_permitted":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.CapPermitted)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.comm":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.Comm
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.container.id":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.ContainerID
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.cookie":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Cookie)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.created_at":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.ProcessContext.Parent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.egid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.EGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.egroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.Credentials.EGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.envs_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return false
				}
				return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.ProcessContext.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.euid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.EUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.euser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.Credentials.EUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Parent.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Parent.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Parent.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Parent.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return false
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Parent.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Parent.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Parent.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Parent.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Parent.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Parent.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.ProcessContext.Parent.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Parent.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.fsgid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.FSGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.fsgroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.Credentials.FSGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.fsuid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.FSUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.fsuser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.Credentials.FSUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.Credentials.Group
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return false
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.interpreter.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.interpreter.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				if !ev.ProcessContext.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "process.parent.is_kworker":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return false
				}
				return ev.ProcessContext.Parent.PIDContext.IsKworker
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.is_thread":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return false
				}
				return ev.ProcessContext.Parent.IsThread
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.pid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.PIDContext.Pid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.ppid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.PPid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.tid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.PIDContext.Tid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.tty_name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.TTYName
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return 0
				}
				return int(ev.ProcessContext.Parent.Credentials.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.parent.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.ProcessContext.HasParent() {
					return ""
				}
				return ev.ProcessContext.Parent.Credentials.User
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.pid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.PIDContext.Pid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.ppid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.PPid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.tid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.PIDContext.Tid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.tty_name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.TTYName
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.ProcessContext.Process.Credentials.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "process.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.ProcessContext.Process.Credentials.User
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.destination.name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveXAttrName(ev, &ev.RemoveXAttr)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.destination.namespace":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveXAttrNamespace(ev, &ev.RemoveXAttr)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.RemoveXAttr.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.RemoveXAttr.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.RemoveXAttr.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.RemoveXAttr.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.RemoveXAttr.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.RemoveXAttr.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.RemoveXAttr.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.RemoveXAttr.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "removexattr.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.RemoveXAttr.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "removexattr.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.RemoveXAttr.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.Old.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.New.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Rename.New)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.New.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Rename.New.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Rename.New.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.New.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.New.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.New.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.New.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.New)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.New))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.New)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.New))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Rename.New.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.destination.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.New.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.destination.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Rename.New.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Rename.Old)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.Old.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Rename.Old.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Rename.Old.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.Old.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.Old.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.Old.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.Old.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.Old)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.Old))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.Old)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.Old))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Rename.Old.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.Old.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rename.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Rename.Old.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rename.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rename.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Rmdir.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Rmdir.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Rmdir.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rmdir.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rmdir.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rmdir.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Rmdir.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Rmdir.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "rmdir.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Rmdir.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "rmdir.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Rmdir.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setgid.egid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetGID.EGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setgid.egroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveSetgidEGroup(ev, &ev.SetGID)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setgid.fsgid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetGID.FSGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setgid.fsgroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveSetgidFSGroup(ev, &ev.SetGID)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setgid.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetGID.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setgid.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveSetgidGroup(ev, &ev.SetGID)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setuid.euid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetUID.EUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setuid.euser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveSetuidEUser(ev, &ev.SetUID)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setuid.fsuid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetUID.FSUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setuid.fsuser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveSetuidFSUser(ev, &ev.SetUID)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setuid.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetUID.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setuid.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveSetuidUser(ev, &ev.SetUID)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setxattr.file.destination.name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveXAttrName(ev, &ev.SetXAttr)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.destination.namespace":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveXAttrNamespace(ev, &ev.SetXAttr)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.SetXAttr.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setxattr.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.SetXAttr.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.SetXAttr.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setxattr.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setxattr.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setxattr.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setxattr.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.SetXAttr.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.SetXAttr.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.SetXAttr.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.SetXAttr.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.SetXAttr.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "setxattr.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.SetXAttr.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "setxattr.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.SetXAttr.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.pid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.PID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.ancestors.args":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgs(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: 100 * eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgsFlags(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgsOptions(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.args_truncated":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgv(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: 100 * eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.argv0":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessArgv0(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: 100 * eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.cap_effective":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.CapEffective)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.cap_permitted":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.CapPermitted)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.comm":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Comm
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.container.id":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.ContainerID
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.cookie":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Cookie)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.created_at":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &element.ProcessContext.Process))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.egid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.EGID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.egroup":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.EGroup
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessEnvp(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessEnvs(ev, &element.ProcessContext.Process)
					results = append(results, result...)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.envs_truncated":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &element.ProcessContext.Process)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.euid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.EUID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.euser":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.EUser
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.change_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.CTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.filesystem":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.gid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.GID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.group":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.in_upper_layer":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, false)
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.inode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.Inode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.mode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.Mode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.modification_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.MTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.mount_id":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.MountID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.name":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.name.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.path":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.path.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.rights":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.FileEvent.FileFields))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.uid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.FileEvent.FileFields.UID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.file.user":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.IsNotKworker() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.fsgid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.FSGID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.fsgroup":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.FSGroup
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.fsuid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.FSUID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.fsuser":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.FSUser
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.gid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.GID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.group":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.Group
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.change_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.filesystem":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.gid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.group":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.in_upper_layer":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, false)
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.inode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.mode":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.modification_time":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.mount_id":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.name":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.name.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.path":
		return &eval.StringArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.path.length":
		return &eval.IntArrayEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.rights":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields))
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.uid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, 0)
						value = iterator.Next()
						continue
					}
					result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.interpreter.file.user":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					if !element.ProcessContext.Process.HasInterpreter() {
						results = append(results, "")
						value = iterator.Next()
						continue
					}
					result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.is_kworker":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.PIDContext.IsKworker
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.is_thread":
		return &eval.BoolArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []bool {
				if result, ok := ctx.BoolCache[field]; ok {
					return result
				}
				var results []bool
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.IsThread
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.BoolCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.pid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.PIDContext.Pid)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.ppid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.PPid)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.tid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.PIDContext.Tid)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.tty_name":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.TTYName
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.uid":
		return &eval.IntArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []int {
				if result, ok := ctx.IntCache[field]; ok {
					return result
				}
				var results []int
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := int(element.ProcessContext.Process.Credentials.UID)
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.IntCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.ancestors.user":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				if result, ok := ctx.StringCache[field]; ok {
					return result
				}
				var results []string
				iterator := &ProcessAncestorsIterator{}
				value := iterator.Front(ctx)
				for value != nil {
					element := (*ProcessCacheEntry)(value)
					result := element.ProcessContext.Process.Credentials.User
					results = append(results, result)
					value = iterator.Next()
				}
				ctx.StringCache[field] = results
				return results
			}, Field: field,
			Weight: eval.IteratorWeight,
		}, nil
	case "signal.target.args":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgs(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "signal.target.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsFlags(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsOptions(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.args_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "signal.target.argv0":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessArgv0(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "signal.target.cap_effective":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.CapEffective)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.cap_permitted":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.CapPermitted)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.comm":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.Comm
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.container.id":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.ContainerID
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.cookie":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Cookie)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.created_at":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &ev.Signal.Target.Process))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.egid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.EGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.egroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.Credentials.EGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvp(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvs(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.envs_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &ev.Signal.Target.Process)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.euid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.EUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.euser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.Credentials.EUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Process.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Process.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Process.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Process.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Process.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Process.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Process.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Process.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Process.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.fsgid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.FSGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.fsgroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.Credentials.FSGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.fsuid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.FSUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.fsuser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.Credentials.FSUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.Credentials.Group
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.interpreter.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.interpreter.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.Process.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.is_kworker":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.PIDContext.IsKworker
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.is_thread":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.IsThread
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.args":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.FieldHandlers.ResolveProcessArgs(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "signal.target.parent.args_flags":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.args_options":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.args_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return false
				}
				return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.argv":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessArgv(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "signal.target.parent.argv0":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: 100 * eval.HandlerWeight,
		}, nil
	case "signal.target.parent.cap_effective":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.CapEffective)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.cap_permitted":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.CapPermitted)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.comm":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.Comm
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.container.id":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.ContainerID
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.cookie":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Cookie)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.created_at":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.Signal.Target.Parent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.egid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.EGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.egroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.Credentials.EGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.envp":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.envs":
		return &eval.StringArrayEvaluator{
			EvalFnc: func(ctx *eval.Context) []string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return []string{}
				}
				return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.envs_truncated":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return false
				}
				return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.Signal.Target.Parent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.euid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.EUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.euser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.Credentials.EUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Parent.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Parent.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Parent.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Parent.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return false
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Parent.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Parent.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Parent.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Parent.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Parent.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Parent.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return 0
				}
				return int(ev.Signal.Target.Parent.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.IsNotKworker() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Parent.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.fsgid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.FSGID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.fsgroup":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.Credentials.FSGroup
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.fsuid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.FSUID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.fsuser":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.Credentials.FSUser
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.Credentials.Group
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return false
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return false
				}
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.interpreter.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return 0
				}
				return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.interpreter.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				if !ev.Signal.Target.Parent.HasInterpreter() {
					return ""
				}
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "signal.target.parent.is_kworker":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return false
				}
				return ev.Signal.Target.Parent.PIDContext.IsKworker
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.is_thread":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return false
				}
				return ev.Signal.Target.Parent.IsThread
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.pid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.PIDContext.Pid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.ppid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.PPid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.tid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.PIDContext.Tid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.tty_name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.TTYName
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return 0
				}
				return int(ev.Signal.Target.Parent.Credentials.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.parent.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				if !ev.Signal.Target.HasParent() {
					return ""
				}
				return ev.Signal.Target.Parent.Credentials.User
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.pid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.PIDContext.Pid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.ppid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.PPid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.tid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.PIDContext.Tid)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.tty_name":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.TTYName
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Target.Process.Credentials.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.target.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.Signal.Target.Process.Credentials.User
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "signal.type":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Signal.Type)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Splice.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Splice.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Splice.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Splice.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Splice.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Splice.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Splice.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Splice.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Splice.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "splice.pipe_entry_flag":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.PipeEntryFlag)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.pipe_exit_flag":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.PipeExitFlag)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "splice.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Splice.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Unlink.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Unlink.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Unlink.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Unlink.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Unlink.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Unlink.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Unlink.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Unlink.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Unlink.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "unlink.flags":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.Flags)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "unlink.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Unlink.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.change_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.File.FileFields.CTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.filesystem":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Utimes.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.gid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.File.FileFields.GID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.group":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Utimes.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.in_upper_layer":
		return &eval.BoolEvaluator{
			EvalFnc: func(ctx *eval.Context) bool {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Utimes.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.inode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.File.FileFields.Inode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.mode":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.File.FileFields.Mode)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.modification_time":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.File.FileFields.MTime)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.mount_id":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.File.FileFields.MountID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.name":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Utimes.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.name.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkBasename,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFileBasename(ev, &ev.Utimes.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.path":
		return &eval.StringEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFilePath(ev, &ev.Utimes.File)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.path.length":
		return &eval.IntEvaluator{
			OpOverrides: ProcessSymlinkPathname,
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return len(ev.FieldHandlers.ResolveFilePath(ev, &ev.Utimes.File))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.rights":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.FieldHandlers.ResolveRights(ev, &ev.Utimes.File.FileFields))
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.file.uid":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.File.FileFields.UID)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	case "utimes.file.user":
		return &eval.StringEvaluator{
			EvalFnc: func(ctx *eval.Context) string {
				ev := ctx.Event.(*Event)
				return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Utimes.File.FileFields)
			},
			Field:  field,
			Weight: eval.HandlerWeight,
		}, nil
	case "utimes.retval":
		return &eval.IntEvaluator{
			EvalFnc: func(ctx *eval.Context) int {
				ev := ctx.Event.(*Event)
				return int(ev.Utimes.SyscallEvent.Retval)
			},
			Field:  field,
			Weight: eval.FunctionWeight,
		}, nil
	}
	return nil, &eval.ErrFieldNotFound{Field: field}
}
func (ev *Event) GetFields() []eval.Field {
	return []eval.Field{
		"async",
		"bind.addr.family",
		"bind.addr.ip",
		"bind.addr.port",
		"bind.retval",
		"capset.cap_effective",
		"capset.cap_permitted",
		"chmod.file.change_time",
		"chmod.file.destination.mode",
		"chmod.file.destination.rights",
		"chmod.file.filesystem",
		"chmod.file.gid",
		"chmod.file.group",
		"chmod.file.in_upper_layer",
		"chmod.file.inode",
		"chmod.file.mode",
		"chmod.file.modification_time",
		"chmod.file.mount_id",
		"chmod.file.name",
		"chmod.file.name.length",
		"chmod.file.path",
		"chmod.file.path.length",
		"chmod.file.rights",
		"chmod.file.uid",
		"chmod.file.user",
		"chmod.retval",
		"chown.file.change_time",
		"chown.file.destination.gid",
		"chown.file.destination.group",
		"chown.file.destination.uid",
		"chown.file.destination.user",
		"chown.file.filesystem",
		"chown.file.gid",
		"chown.file.group",
		"chown.file.in_upper_layer",
		"chown.file.inode",
		"chown.file.mode",
		"chown.file.modification_time",
		"chown.file.mount_id",
		"chown.file.name",
		"chown.file.name.length",
		"chown.file.path",
		"chown.file.path.length",
		"chown.file.rights",
		"chown.file.uid",
		"chown.file.user",
		"chown.retval",
		"container.id",
		"container.tags",
		"dns.id",
		"dns.question.class",
		"dns.question.count",
		"dns.question.length",
		"dns.question.name",
		"dns.question.type",
		"exec.args",
		"exec.args_flags",
		"exec.args_options",
		"exec.args_truncated",
		"exec.argv",
		"exec.argv0",
		"exec.cap_effective",
		"exec.cap_permitted",
		"exec.comm",
		"exec.container.id",
		"exec.cookie",
		"exec.created_at",
		"exec.egid",
		"exec.egroup",
		"exec.envp",
		"exec.envs",
		"exec.envs_truncated",
		"exec.euid",
		"exec.euser",
		"exec.file.change_time",
		"exec.file.filesystem",
		"exec.file.gid",
		"exec.file.group",
		"exec.file.in_upper_layer",
		"exec.file.inode",
		"exec.file.mode",
		"exec.file.modification_time",
		"exec.file.mount_id",
		"exec.file.name",
		"exec.file.name.length",
		"exec.file.path",
		"exec.file.path.length",
		"exec.file.rights",
		"exec.file.uid",
		"exec.file.user",
		"exec.fsgid",
		"exec.fsgroup",
		"exec.fsuid",
		"exec.fsuser",
		"exec.gid",
		"exec.group",
		"exec.interpreter.file.change_time",
		"exec.interpreter.file.filesystem",
		"exec.interpreter.file.gid",
		"exec.interpreter.file.group",
		"exec.interpreter.file.in_upper_layer",
		"exec.interpreter.file.inode",
		"exec.interpreter.file.mode",
		"exec.interpreter.file.modification_time",
		"exec.interpreter.file.mount_id",
		"exec.interpreter.file.name",
		"exec.interpreter.file.name.length",
		"exec.interpreter.file.path",
		"exec.interpreter.file.path.length",
		"exec.interpreter.file.rights",
		"exec.interpreter.file.uid",
		"exec.interpreter.file.user",
		"exec.is_kworker",
		"exec.is_thread",
		"exec.pid",
		"exec.ppid",
		"exec.tid",
		"exec.tty_name",
		"exec.uid",
		"exec.user",
		"exit.args",
		"exit.args_flags",
		"exit.args_options",
		"exit.args_truncated",
		"exit.argv",
		"exit.argv0",
		"exit.cap_effective",
		"exit.cap_permitted",
		"exit.cause",
		"exit.code",
		"exit.comm",
		"exit.container.id",
		"exit.cookie",
		"exit.created_at",
		"exit.egid",
		"exit.egroup",
		"exit.envp",
		"exit.envs",
		"exit.envs_truncated",
		"exit.euid",
		"exit.euser",
		"exit.file.change_time",
		"exit.file.filesystem",
		"exit.file.gid",
		"exit.file.group",
		"exit.file.in_upper_layer",
		"exit.file.inode",
		"exit.file.mode",
		"exit.file.modification_time",
		"exit.file.mount_id",
		"exit.file.name",
		"exit.file.name.length",
		"exit.file.path",
		"exit.file.path.length",
		"exit.file.rights",
		"exit.file.uid",
		"exit.file.user",
		"exit.fsgid",
		"exit.fsgroup",
		"exit.fsuid",
		"exit.fsuser",
		"exit.gid",
		"exit.group",
		"exit.interpreter.file.change_time",
		"exit.interpreter.file.filesystem",
		"exit.interpreter.file.gid",
		"exit.interpreter.file.group",
		"exit.interpreter.file.in_upper_layer",
		"exit.interpreter.file.inode",
		"exit.interpreter.file.mode",
		"exit.interpreter.file.modification_time",
		"exit.interpreter.file.mount_id",
		"exit.interpreter.file.name",
		"exit.interpreter.file.name.length",
		"exit.interpreter.file.path",
		"exit.interpreter.file.path.length",
		"exit.interpreter.file.rights",
		"exit.interpreter.file.uid",
		"exit.interpreter.file.user",
		"exit.is_kworker",
		"exit.is_thread",
		"exit.pid",
		"exit.ppid",
		"exit.tid",
		"exit.tty_name",
		"exit.uid",
		"exit.user",
		"link.file.change_time",
		"link.file.destination.change_time",
		"link.file.destination.filesystem",
		"link.file.destination.gid",
		"link.file.destination.group",
		"link.file.destination.in_upper_layer",
		"link.file.destination.inode",
		"link.file.destination.mode",
		"link.file.destination.modification_time",
		"link.file.destination.mount_id",
		"link.file.destination.name",
		"link.file.destination.name.length",
		"link.file.destination.path",
		"link.file.destination.path.length",
		"link.file.destination.rights",
		"link.file.destination.uid",
		"link.file.destination.user",
		"link.file.filesystem",
		"link.file.gid",
		"link.file.group",
		"link.file.in_upper_layer",
		"link.file.inode",
		"link.file.mode",
		"link.file.modification_time",
		"link.file.mount_id",
		"link.file.name",
		"link.file.name.length",
		"link.file.path",
		"link.file.path.length",
		"link.file.rights",
		"link.file.uid",
		"link.file.user",
		"link.retval",
		"mkdir.file.change_time",
		"mkdir.file.destination.mode",
		"mkdir.file.destination.rights",
		"mkdir.file.filesystem",
		"mkdir.file.gid",
		"mkdir.file.group",
		"mkdir.file.in_upper_layer",
		"mkdir.file.inode",
		"mkdir.file.mode",
		"mkdir.file.modification_time",
		"mkdir.file.mount_id",
		"mkdir.file.name",
		"mkdir.file.name.length",
		"mkdir.file.path",
		"mkdir.file.path.length",
		"mkdir.file.rights",
		"mkdir.file.uid",
		"mkdir.file.user",
		"mkdir.retval",
		"mount.fs_type",
		"mount.mountpoint.path",
		"mount.retval",
		"mount.source.path",
		"network.destination.ip",
		"network.destination.port",
		"network.device.ifindex",
		"network.device.ifname",
		"network.l3_protocol",
		"network.l4_protocol",
		"network.size",
		"network.source.ip",
		"network.source.port",
		"open.file.change_time",
		"open.file.destination.mode",
		"open.file.filesystem",
		"open.file.gid",
		"open.file.group",
		"open.file.in_upper_layer",
		"open.file.inode",
		"open.file.mode",
		"open.file.modification_time",
		"open.file.mount_id",
		"open.file.name",
		"open.file.name.length",
		"open.file.path",
		"open.file.path.length",
		"open.file.rights",
		"open.file.uid",
		"open.file.user",
		"open.flags",
		"open.retval",
		"process.ancestors.args",
		"process.ancestors.args_flags",
		"process.ancestors.args_options",
		"process.ancestors.args_truncated",
		"process.ancestors.argv",
		"process.ancestors.argv0",
		"process.ancestors.cap_effective",
		"process.ancestors.cap_permitted",
		"process.ancestors.comm",
		"process.ancestors.container.id",
		"process.ancestors.cookie",
		"process.ancestors.created_at",
		"process.ancestors.egid",
		"process.ancestors.egroup",
		"process.ancestors.envp",
		"process.ancestors.envs",
		"process.ancestors.envs_truncated",
		"process.ancestors.euid",
		"process.ancestors.euser",
		"process.ancestors.file.change_time",
		"process.ancestors.file.filesystem",
		"process.ancestors.file.gid",
		"process.ancestors.file.group",
		"process.ancestors.file.in_upper_layer",
		"process.ancestors.file.inode",
		"process.ancestors.file.mode",
		"process.ancestors.file.modification_time",
		"process.ancestors.file.mount_id",
		"process.ancestors.file.name",
		"process.ancestors.file.name.length",
		"process.ancestors.file.path",
		"process.ancestors.file.path.length",
		"process.ancestors.file.rights",
		"process.ancestors.file.uid",
		"process.ancestors.file.user",
		"process.ancestors.fsgid",
		"process.ancestors.fsgroup",
		"process.ancestors.fsuid",
		"process.ancestors.fsuser",
		"process.ancestors.gid",
		"process.ancestors.group",
		"process.ancestors.interpreter.file.change_time",
		"process.ancestors.interpreter.file.filesystem",
		"process.ancestors.interpreter.file.gid",
		"process.ancestors.interpreter.file.group",
		"process.ancestors.interpreter.file.in_upper_layer",
		"process.ancestors.interpreter.file.inode",
		"process.ancestors.interpreter.file.mode",
		"process.ancestors.interpreter.file.modification_time",
		"process.ancestors.interpreter.file.mount_id",
		"process.ancestors.interpreter.file.name",
		"process.ancestors.interpreter.file.name.length",
		"process.ancestors.interpreter.file.path",
		"process.ancestors.interpreter.file.path.length",
		"process.ancestors.interpreter.file.rights",
		"process.ancestors.interpreter.file.uid",
		"process.ancestors.interpreter.file.user",
		"process.ancestors.is_kworker",
		"process.ancestors.is_thread",
		"process.ancestors.pid",
		"process.ancestors.ppid",
		"process.ancestors.tid",
		"process.ancestors.tty_name",
		"process.ancestors.uid",
		"process.ancestors.user",
		"process.args",
		"process.args_flags",
		"process.args_options",
		"process.args_truncated",
		"process.argv",
		"process.argv0",
		"process.cap_effective",
		"process.cap_permitted",
		"process.comm",
		"process.container.id",
		"process.cookie",
		"process.created_at",
		"process.egid",
		"process.egroup",
		"process.envp",
		"process.envs",
		"process.envs_truncated",
		"process.euid",
		"process.euser",
		"process.file.change_time",
		"process.file.filesystem",
		"process.file.gid",
		"process.file.group",
		"process.file.in_upper_layer",
		"process.file.inode",
		"process.file.mode",
		"process.file.modification_time",
		"process.file.mount_id",
		"process.file.name",
		"process.file.name.length",
		"process.file.path",
		"process.file.path.length",
		"process.file.rights",
		"process.file.uid",
		"process.file.user",
		"process.fsgid",
		"process.fsgroup",
		"process.fsuid",
		"process.fsuser",
		"process.gid",
		"process.group",
		"process.interpreter.file.change_time",
		"process.interpreter.file.filesystem",
		"process.interpreter.file.gid",
		"process.interpreter.file.group",
		"process.interpreter.file.in_upper_layer",
		"process.interpreter.file.inode",
		"process.interpreter.file.mode",
		"process.interpreter.file.modification_time",
		"process.interpreter.file.mount_id",
		"process.interpreter.file.name",
		"process.interpreter.file.name.length",
		"process.interpreter.file.path",
		"process.interpreter.file.path.length",
		"process.interpreter.file.rights",
		"process.interpreter.file.uid",
		"process.interpreter.file.user",
		"process.is_kworker",
		"process.is_thread",
		"process.parent.args",
		"process.parent.args_flags",
		"process.parent.args_options",
		"process.parent.args_truncated",
		"process.parent.argv",
		"process.parent.argv0",
		"process.parent.cap_effective",
		"process.parent.cap_permitted",
		"process.parent.comm",
		"process.parent.container.id",
		"process.parent.cookie",
		"process.parent.created_at",
		"process.parent.egid",
		"process.parent.egroup",
		"process.parent.envp",
		"process.parent.envs",
		"process.parent.envs_truncated",
		"process.parent.euid",
		"process.parent.euser",
		"process.parent.file.change_time",
		"process.parent.file.filesystem",
		"process.parent.file.gid",
		"process.parent.file.group",
		"process.parent.file.in_upper_layer",
		"process.parent.file.inode",
		"process.parent.file.mode",
		"process.parent.file.modification_time",
		"process.parent.file.mount_id",
		"process.parent.file.name",
		"process.parent.file.name.length",
		"process.parent.file.path",
		"process.parent.file.path.length",
		"process.parent.file.rights",
		"process.parent.file.uid",
		"process.parent.file.user",
		"process.parent.fsgid",
		"process.parent.fsgroup",
		"process.parent.fsuid",
		"process.parent.fsuser",
		"process.parent.gid",
		"process.parent.group",
		"process.parent.interpreter.file.change_time",
		"process.parent.interpreter.file.filesystem",
		"process.parent.interpreter.file.gid",
		"process.parent.interpreter.file.group",
		"process.parent.interpreter.file.in_upper_layer",
		"process.parent.interpreter.file.inode",
		"process.parent.interpreter.file.mode",
		"process.parent.interpreter.file.modification_time",
		"process.parent.interpreter.file.mount_id",
		"process.parent.interpreter.file.name",
		"process.parent.interpreter.file.name.length",
		"process.parent.interpreter.file.path",
		"process.parent.interpreter.file.path.length",
		"process.parent.interpreter.file.rights",
		"process.parent.interpreter.file.uid",
		"process.parent.interpreter.file.user",
		"process.parent.is_kworker",
		"process.parent.is_thread",
		"process.parent.pid",
		"process.parent.ppid",
		"process.parent.tid",
		"process.parent.tty_name",
		"process.parent.uid",
		"process.parent.user",
		"process.pid",
		"process.ppid",
		"process.tid",
		"process.tty_name",
		"process.uid",
		"process.user",
		"removexattr.file.change_time",
		"removexattr.file.destination.name",
		"removexattr.file.destination.namespace",
		"removexattr.file.filesystem",
		"removexattr.file.gid",
		"removexattr.file.group",
		"removexattr.file.in_upper_layer",
		"removexattr.file.inode",
		"removexattr.file.mode",
		"removexattr.file.modification_time",
		"removexattr.file.mount_id",
		"removexattr.file.name",
		"removexattr.file.name.length",
		"removexattr.file.path",
		"removexattr.file.path.length",
		"removexattr.file.rights",
		"removexattr.file.uid",
		"removexattr.file.user",
		"removexattr.retval",
		"rename.file.change_time",
		"rename.file.destination.change_time",
		"rename.file.destination.filesystem",
		"rename.file.destination.gid",
		"rename.file.destination.group",
		"rename.file.destination.in_upper_layer",
		"rename.file.destination.inode",
		"rename.file.destination.mode",
		"rename.file.destination.modification_time",
		"rename.file.destination.mount_id",
		"rename.file.destination.name",
		"rename.file.destination.name.length",
		"rename.file.destination.path",
		"rename.file.destination.path.length",
		"rename.file.destination.rights",
		"rename.file.destination.uid",
		"rename.file.destination.user",
		"rename.file.filesystem",
		"rename.file.gid",
		"rename.file.group",
		"rename.file.in_upper_layer",
		"rename.file.inode",
		"rename.file.mode",
		"rename.file.modification_time",
		"rename.file.mount_id",
		"rename.file.name",
		"rename.file.name.length",
		"rename.file.path",
		"rename.file.path.length",
		"rename.file.rights",
		"rename.file.uid",
		"rename.file.user",
		"rename.retval",
		"rmdir.file.change_time",
		"rmdir.file.filesystem",
		"rmdir.file.gid",
		"rmdir.file.group",
		"rmdir.file.in_upper_layer",
		"rmdir.file.inode",
		"rmdir.file.mode",
		"rmdir.file.modification_time",
		"rmdir.file.mount_id",
		"rmdir.file.name",
		"rmdir.file.name.length",
		"rmdir.file.path",
		"rmdir.file.path.length",
		"rmdir.file.rights",
		"rmdir.file.uid",
		"rmdir.file.user",
		"rmdir.retval",
		"setgid.egid",
		"setgid.egroup",
		"setgid.fsgid",
		"setgid.fsgroup",
		"setgid.gid",
		"setgid.group",
		"setuid.euid",
		"setuid.euser",
		"setuid.fsuid",
		"setuid.fsuser",
		"setuid.uid",
		"setuid.user",
		"setxattr.file.change_time",
		"setxattr.file.destination.name",
		"setxattr.file.destination.namespace",
		"setxattr.file.filesystem",
		"setxattr.file.gid",
		"setxattr.file.group",
		"setxattr.file.in_upper_layer",
		"setxattr.file.inode",
		"setxattr.file.mode",
		"setxattr.file.modification_time",
		"setxattr.file.mount_id",
		"setxattr.file.name",
		"setxattr.file.name.length",
		"setxattr.file.path",
		"setxattr.file.path.length",
		"setxattr.file.rights",
		"setxattr.file.uid",
		"setxattr.file.user",
		"setxattr.retval",
		"signal.pid",
		"signal.retval",
		"signal.target.ancestors.args",
		"signal.target.ancestors.args_flags",
		"signal.target.ancestors.args_options",
		"signal.target.ancestors.args_truncated",
		"signal.target.ancestors.argv",
		"signal.target.ancestors.argv0",
		"signal.target.ancestors.cap_effective",
		"signal.target.ancestors.cap_permitted",
		"signal.target.ancestors.comm",
		"signal.target.ancestors.container.id",
		"signal.target.ancestors.cookie",
		"signal.target.ancestors.created_at",
		"signal.target.ancestors.egid",
		"signal.target.ancestors.egroup",
		"signal.target.ancestors.envp",
		"signal.target.ancestors.envs",
		"signal.target.ancestors.envs_truncated",
		"signal.target.ancestors.euid",
		"signal.target.ancestors.euser",
		"signal.target.ancestors.file.change_time",
		"signal.target.ancestors.file.filesystem",
		"signal.target.ancestors.file.gid",
		"signal.target.ancestors.file.group",
		"signal.target.ancestors.file.in_upper_layer",
		"signal.target.ancestors.file.inode",
		"signal.target.ancestors.file.mode",
		"signal.target.ancestors.file.modification_time",
		"signal.target.ancestors.file.mount_id",
		"signal.target.ancestors.file.name",
		"signal.target.ancestors.file.name.length",
		"signal.target.ancestors.file.path",
		"signal.target.ancestors.file.path.length",
		"signal.target.ancestors.file.rights",
		"signal.target.ancestors.file.uid",
		"signal.target.ancestors.file.user",
		"signal.target.ancestors.fsgid",
		"signal.target.ancestors.fsgroup",
		"signal.target.ancestors.fsuid",
		"signal.target.ancestors.fsuser",
		"signal.target.ancestors.gid",
		"signal.target.ancestors.group",
		"signal.target.ancestors.interpreter.file.change_time",
		"signal.target.ancestors.interpreter.file.filesystem",
		"signal.target.ancestors.interpreter.file.gid",
		"signal.target.ancestors.interpreter.file.group",
		"signal.target.ancestors.interpreter.file.in_upper_layer",
		"signal.target.ancestors.interpreter.file.inode",
		"signal.target.ancestors.interpreter.file.mode",
		"signal.target.ancestors.interpreter.file.modification_time",
		"signal.target.ancestors.interpreter.file.mount_id",
		"signal.target.ancestors.interpreter.file.name",
		"signal.target.ancestors.interpreter.file.name.length",
		"signal.target.ancestors.interpreter.file.path",
		"signal.target.ancestors.interpreter.file.path.length",
		"signal.target.ancestors.interpreter.file.rights",
		"signal.target.ancestors.interpreter.file.uid",
		"signal.target.ancestors.interpreter.file.user",
		"signal.target.ancestors.is_kworker",
		"signal.target.ancestors.is_thread",
		"signal.target.ancestors.pid",
		"signal.target.ancestors.ppid",
		"signal.target.ancestors.tid",
		"signal.target.ancestors.tty_name",
		"signal.target.ancestors.uid",
		"signal.target.ancestors.user",
		"signal.target.args",
		"signal.target.args_flags",
		"signal.target.args_options",
		"signal.target.args_truncated",
		"signal.target.argv",
		"signal.target.argv0",
		"signal.target.cap_effective",
		"signal.target.cap_permitted",
		"signal.target.comm",
		"signal.target.container.id",
		"signal.target.cookie",
		"signal.target.created_at",
		"signal.target.egid",
		"signal.target.egroup",
		"signal.target.envp",
		"signal.target.envs",
		"signal.target.envs_truncated",
		"signal.target.euid",
		"signal.target.euser",
		"signal.target.file.change_time",
		"signal.target.file.filesystem",
		"signal.target.file.gid",
		"signal.target.file.group",
		"signal.target.file.in_upper_layer",
		"signal.target.file.inode",
		"signal.target.file.mode",
		"signal.target.file.modification_time",
		"signal.target.file.mount_id",
		"signal.target.file.name",
		"signal.target.file.name.length",
		"signal.target.file.path",
		"signal.target.file.path.length",
		"signal.target.file.rights",
		"signal.target.file.uid",
		"signal.target.file.user",
		"signal.target.fsgid",
		"signal.target.fsgroup",
		"signal.target.fsuid",
		"signal.target.fsuser",
		"signal.target.gid",
		"signal.target.group",
		"signal.target.interpreter.file.change_time",
		"signal.target.interpreter.file.filesystem",
		"signal.target.interpreter.file.gid",
		"signal.target.interpreter.file.group",
		"signal.target.interpreter.file.in_upper_layer",
		"signal.target.interpreter.file.inode",
		"signal.target.interpreter.file.mode",
		"signal.target.interpreter.file.modification_time",
		"signal.target.interpreter.file.mount_id",
		"signal.target.interpreter.file.name",
		"signal.target.interpreter.file.name.length",
		"signal.target.interpreter.file.path",
		"signal.target.interpreter.file.path.length",
		"signal.target.interpreter.file.rights",
		"signal.target.interpreter.file.uid",
		"signal.target.interpreter.file.user",
		"signal.target.is_kworker",
		"signal.target.is_thread",
		"signal.target.parent.args",
		"signal.target.parent.args_flags",
		"signal.target.parent.args_options",
		"signal.target.parent.args_truncated",
		"signal.target.parent.argv",
		"signal.target.parent.argv0",
		"signal.target.parent.cap_effective",
		"signal.target.parent.cap_permitted",
		"signal.target.parent.comm",
		"signal.target.parent.container.id",
		"signal.target.parent.cookie",
		"signal.target.parent.created_at",
		"signal.target.parent.egid",
		"signal.target.parent.egroup",
		"signal.target.parent.envp",
		"signal.target.parent.envs",
		"signal.target.parent.envs_truncated",
		"signal.target.parent.euid",
		"signal.target.parent.euser",
		"signal.target.parent.file.change_time",
		"signal.target.parent.file.filesystem",
		"signal.target.parent.file.gid",
		"signal.target.parent.file.group",
		"signal.target.parent.file.in_upper_layer",
		"signal.target.parent.file.inode",
		"signal.target.parent.file.mode",
		"signal.target.parent.file.modification_time",
		"signal.target.parent.file.mount_id",
		"signal.target.parent.file.name",
		"signal.target.parent.file.name.length",
		"signal.target.parent.file.path",
		"signal.target.parent.file.path.length",
		"signal.target.parent.file.rights",
		"signal.target.parent.file.uid",
		"signal.target.parent.file.user",
		"signal.target.parent.fsgid",
		"signal.target.parent.fsgroup",
		"signal.target.parent.fsuid",
		"signal.target.parent.fsuser",
		"signal.target.parent.gid",
		"signal.target.parent.group",
		"signal.target.parent.interpreter.file.change_time",
		"signal.target.parent.interpreter.file.filesystem",
		"signal.target.parent.interpreter.file.gid",
		"signal.target.parent.interpreter.file.group",
		"signal.target.parent.interpreter.file.in_upper_layer",
		"signal.target.parent.interpreter.file.inode",
		"signal.target.parent.interpreter.file.mode",
		"signal.target.parent.interpreter.file.modification_time",
		"signal.target.parent.interpreter.file.mount_id",
		"signal.target.parent.interpreter.file.name",
		"signal.target.parent.interpreter.file.name.length",
		"signal.target.parent.interpreter.file.path",
		"signal.target.parent.interpreter.file.path.length",
		"signal.target.parent.interpreter.file.rights",
		"signal.target.parent.interpreter.file.uid",
		"signal.target.parent.interpreter.file.user",
		"signal.target.parent.is_kworker",
		"signal.target.parent.is_thread",
		"signal.target.parent.pid",
		"signal.target.parent.ppid",
		"signal.target.parent.tid",
		"signal.target.parent.tty_name",
		"signal.target.parent.uid",
		"signal.target.parent.user",
		"signal.target.pid",
		"signal.target.ppid",
		"signal.target.tid",
		"signal.target.tty_name",
		"signal.target.uid",
		"signal.target.user",
		"signal.type",
		"splice.file.change_time",
		"splice.file.filesystem",
		"splice.file.gid",
		"splice.file.group",
		"splice.file.in_upper_layer",
		"splice.file.inode",
		"splice.file.mode",
		"splice.file.modification_time",
		"splice.file.mount_id",
		"splice.file.name",
		"splice.file.name.length",
		"splice.file.path",
		"splice.file.path.length",
		"splice.file.rights",
		"splice.file.uid",
		"splice.file.user",
		"splice.pipe_entry_flag",
		"splice.pipe_exit_flag",
		"splice.retval",
		"unlink.file.change_time",
		"unlink.file.filesystem",
		"unlink.file.gid",
		"unlink.file.group",
		"unlink.file.in_upper_layer",
		"unlink.file.inode",
		"unlink.file.mode",
		"unlink.file.modification_time",
		"unlink.file.mount_id",
		"unlink.file.name",
		"unlink.file.name.length",
		"unlink.file.path",
		"unlink.file.path.length",
		"unlink.file.rights",
		"unlink.file.uid",
		"unlink.file.user",
		"unlink.flags",
		"unlink.retval",
		"utimes.file.change_time",
		"utimes.file.filesystem",
		"utimes.file.gid",
		"utimes.file.group",
		"utimes.file.in_upper_layer",
		"utimes.file.inode",
		"utimes.file.mode",
		"utimes.file.modification_time",
		"utimes.file.mount_id",
		"utimes.file.name",
		"utimes.file.name.length",
		"utimes.file.path",
		"utimes.file.path.length",
		"utimes.file.rights",
		"utimes.file.uid",
		"utimes.file.user",
		"utimes.retval",
	}
}
func (ev *Event) GetFieldValue(field eval.Field) (interface{}, error) {
	switch field {
	case "async":
		return ev.Async, nil
	case "bind.addr.family":
		return int(ev.Bind.AddrFamily), nil
	case "bind.addr.ip":
		return ev.Bind.Addr.IPNet, nil
	case "bind.addr.port":
		return int(ev.Bind.Addr.Port), nil
	case "bind.retval":
		return int(ev.Bind.SyscallEvent.Retval), nil
	case "capset.cap_effective":
		return int(ev.Capset.CapEffective), nil
	case "capset.cap_permitted":
		return int(ev.Capset.CapPermitted), nil
	case "chmod.file.change_time":
		return int(ev.Chmod.File.FileFields.CTime), nil
	case "chmod.file.destination.mode":
		return int(ev.Chmod.Mode), nil
	case "chmod.file.destination.rights":
		return int(ev.Chmod.Mode), nil
	case "chmod.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Chmod.File), nil
	case "chmod.file.gid":
		return int(ev.Chmod.File.FileFields.GID), nil
	case "chmod.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Chmod.File.FileFields), nil
	case "chmod.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Chmod.File.FileFields), nil
	case "chmod.file.inode":
		return int(ev.Chmod.File.FileFields.Inode), nil
	case "chmod.file.mode":
		return int(ev.Chmod.File.FileFields.Mode), nil
	case "chmod.file.modification_time":
		return int(ev.Chmod.File.FileFields.MTime), nil
	case "chmod.file.mount_id":
		return int(ev.Chmod.File.FileFields.MountID), nil
	case "chmod.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chmod.File), nil
	case "chmod.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chmod.File), nil
	case "chmod.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Chmod.File), nil
	case "chmod.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Chmod.File), nil
	case "chmod.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Chmod.File.FileFields)), nil
	case "chmod.file.uid":
		return int(ev.Chmod.File.FileFields.UID), nil
	case "chmod.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Chmod.File.FileFields), nil
	case "chmod.retval":
		return int(ev.Chmod.SyscallEvent.Retval), nil
	case "chown.file.change_time":
		return int(ev.Chown.File.FileFields.CTime), nil
	case "chown.file.destination.gid":
		return int(ev.Chown.GID), nil
	case "chown.file.destination.group":
		return ev.FieldHandlers.ResolveChownGID(ev, &ev.Chown), nil
	case "chown.file.destination.uid":
		return int(ev.Chown.UID), nil
	case "chown.file.destination.user":
		return ev.FieldHandlers.ResolveChownUID(ev, &ev.Chown), nil
	case "chown.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Chown.File), nil
	case "chown.file.gid":
		return int(ev.Chown.File.FileFields.GID), nil
	case "chown.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Chown.File.FileFields), nil
	case "chown.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Chown.File.FileFields), nil
	case "chown.file.inode":
		return int(ev.Chown.File.FileFields.Inode), nil
	case "chown.file.mode":
		return int(ev.Chown.File.FileFields.Mode), nil
	case "chown.file.modification_time":
		return int(ev.Chown.File.FileFields.MTime), nil
	case "chown.file.mount_id":
		return int(ev.Chown.File.FileFields.MountID), nil
	case "chown.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chown.File), nil
	case "chown.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Chown.File), nil
	case "chown.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Chown.File), nil
	case "chown.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Chown.File), nil
	case "chown.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Chown.File.FileFields)), nil
	case "chown.file.uid":
		return int(ev.Chown.File.FileFields.UID), nil
	case "chown.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Chown.File.FileFields), nil
	case "chown.retval":
		return int(ev.Chown.SyscallEvent.Retval), nil
	case "container.id":
		return ev.FieldHandlers.ResolveContainerID(ev, &ev.ContainerContext), nil
	case "container.tags":
		return ev.FieldHandlers.ResolveContainerTags(ev, &ev.ContainerContext), nil
	case "dns.id":
		return int(ev.DNS.ID), nil
	case "dns.question.class":
		return int(ev.DNS.Class), nil
	case "dns.question.count":
		return int(ev.DNS.Count), nil
	case "dns.question.length":
		return int(ev.DNS.Size), nil
	case "dns.question.name":
		return len(ev.DNS.Name), nil
	case "dns.question.type":
		return int(ev.DNS.Type), nil
	case "exec.args":
		return ev.FieldHandlers.ResolveProcessArgs(ev, ev.Exec.Process), nil
	case "exec.args_flags":
		return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.Exec.Process), nil
	case "exec.args_options":
		return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.Exec.Process), nil
	case "exec.args_truncated":
		return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.Exec.Process), nil
	case "exec.argv":
		return ev.FieldHandlers.ResolveProcessArgv(ev, ev.Exec.Process), nil
	case "exec.argv0":
		return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.Exec.Process), nil
	case "exec.cap_effective":
		return int(ev.Exec.Process.Credentials.CapEffective), nil
	case "exec.cap_permitted":
		return int(ev.Exec.Process.Credentials.CapPermitted), nil
	case "exec.comm":
		return ev.Exec.Process.Comm, nil
	case "exec.container.id":
		return ev.Exec.Process.ContainerID, nil
	case "exec.cookie":
		return int(ev.Exec.Process.Cookie), nil
	case "exec.created_at":
		return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.Exec.Process)), nil
	case "exec.egid":
		return int(ev.Exec.Process.Credentials.EGID), nil
	case "exec.egroup":
		return ev.Exec.Process.Credentials.EGroup, nil
	case "exec.envp":
		return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.Exec.Process), nil
	case "exec.envs":
		return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.Exec.Process), nil
	case "exec.envs_truncated":
		return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.Exec.Process), nil
	case "exec.euid":
		return int(ev.Exec.Process.Credentials.EUID), nil
	case "exec.euser":
		return ev.Exec.Process.Credentials.EUser, nil
	case "exec.file.change_time":
		return int(ev.Exec.Process.FileEvent.FileFields.CTime), nil
	case "exec.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exec.Process.FileEvent), nil
	case "exec.file.gid":
		return int(ev.Exec.Process.FileEvent.FileFields.GID), nil
	case "exec.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exec.Process.FileEvent.FileFields), nil
	case "exec.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exec.Process.FileEvent.FileFields), nil
	case "exec.file.inode":
		return int(ev.Exec.Process.FileEvent.FileFields.Inode), nil
	case "exec.file.mode":
		return int(ev.Exec.Process.FileEvent.FileFields.Mode), nil
	case "exec.file.modification_time":
		return int(ev.Exec.Process.FileEvent.FileFields.MTime), nil
	case "exec.file.mount_id":
		return int(ev.Exec.Process.FileEvent.FileFields.MountID), nil
	case "exec.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.FileEvent), nil
	case "exec.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.FileEvent), nil
	case "exec.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.FileEvent), nil
	case "exec.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.FileEvent), nil
	case "exec.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exec.Process.FileEvent.FileFields)), nil
	case "exec.file.uid":
		return int(ev.Exec.Process.FileEvent.FileFields.UID), nil
	case "exec.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exec.Process.FileEvent.FileFields), nil
	case "exec.fsgid":
		return int(ev.Exec.Process.Credentials.FSGID), nil
	case "exec.fsgroup":
		return ev.Exec.Process.Credentials.FSGroup, nil
	case "exec.fsuid":
		return int(ev.Exec.Process.Credentials.FSUID), nil
	case "exec.fsuser":
		return ev.Exec.Process.Credentials.FSUser, nil
	case "exec.gid":
		return int(ev.Exec.Process.Credentials.GID), nil
	case "exec.group":
		return ev.Exec.Process.Credentials.Group, nil
	case "exec.interpreter.file.change_time":
		return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.CTime), nil
	case "exec.interpreter.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exec.Process.LinuxBinprm.FileEvent), nil
	case "exec.interpreter.file.gid":
		return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.GID), nil
	case "exec.interpreter.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "exec.interpreter.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "exec.interpreter.file.inode":
		return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Inode), nil
	case "exec.interpreter.file.mode":
		return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Mode), nil
	case "exec.interpreter.file.modification_time":
		return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.MTime), nil
	case "exec.interpreter.file.mount_id":
		return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.MountID), nil
	case "exec.interpreter.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.LinuxBinprm.FileEvent), nil
	case "exec.interpreter.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exec.Process.LinuxBinprm.FileEvent), nil
	case "exec.interpreter.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.LinuxBinprm.FileEvent), nil
	case "exec.interpreter.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exec.Process.LinuxBinprm.FileEvent), nil
	case "exec.interpreter.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields)), nil
	case "exec.interpreter.file.uid":
		return int(ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.UID), nil
	case "exec.interpreter.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exec.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "exec.is_kworker":
		return ev.Exec.Process.PIDContext.IsKworker, nil
	case "exec.is_thread":
		return ev.Exec.Process.IsThread, nil
	case "exec.pid":
		return int(ev.Exec.Process.PIDContext.Pid), nil
	case "exec.ppid":
		return int(ev.Exec.Process.PPid), nil
	case "exec.tid":
		return int(ev.Exec.Process.PIDContext.Tid), nil
	case "exec.tty_name":
		return ev.Exec.Process.TTYName, nil
	case "exec.uid":
		return int(ev.Exec.Process.Credentials.UID), nil
	case "exec.user":
		return ev.Exec.Process.Credentials.User, nil
	case "exit.args":
		return ev.FieldHandlers.ResolveProcessArgs(ev, ev.Exit.Process), nil
	case "exit.args_flags":
		return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.Exit.Process), nil
	case "exit.args_options":
		return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.Exit.Process), nil
	case "exit.args_truncated":
		return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.Exit.Process), nil
	case "exit.argv":
		return ev.FieldHandlers.ResolveProcessArgv(ev, ev.Exit.Process), nil
	case "exit.argv0":
		return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.Exit.Process), nil
	case "exit.cap_effective":
		return int(ev.Exit.Process.Credentials.CapEffective), nil
	case "exit.cap_permitted":
		return int(ev.Exit.Process.Credentials.CapPermitted), nil
	case "exit.cause":
		return int(ev.Exit.Cause), nil
	case "exit.code":
		return int(ev.Exit.Code), nil
	case "exit.comm":
		return ev.Exit.Process.Comm, nil
	case "exit.container.id":
		return ev.Exit.Process.ContainerID, nil
	case "exit.cookie":
		return int(ev.Exit.Process.Cookie), nil
	case "exit.created_at":
		return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.Exit.Process)), nil
	case "exit.egid":
		return int(ev.Exit.Process.Credentials.EGID), nil
	case "exit.egroup":
		return ev.Exit.Process.Credentials.EGroup, nil
	case "exit.envp":
		return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.Exit.Process), nil
	case "exit.envs":
		return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.Exit.Process), nil
	case "exit.envs_truncated":
		return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.Exit.Process), nil
	case "exit.euid":
		return int(ev.Exit.Process.Credentials.EUID), nil
	case "exit.euser":
		return ev.Exit.Process.Credentials.EUser, nil
	case "exit.file.change_time":
		return int(ev.Exit.Process.FileEvent.FileFields.CTime), nil
	case "exit.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exit.Process.FileEvent), nil
	case "exit.file.gid":
		return int(ev.Exit.Process.FileEvent.FileFields.GID), nil
	case "exit.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exit.Process.FileEvent.FileFields), nil
	case "exit.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exit.Process.FileEvent.FileFields), nil
	case "exit.file.inode":
		return int(ev.Exit.Process.FileEvent.FileFields.Inode), nil
	case "exit.file.mode":
		return int(ev.Exit.Process.FileEvent.FileFields.Mode), nil
	case "exit.file.modification_time":
		return int(ev.Exit.Process.FileEvent.FileFields.MTime), nil
	case "exit.file.mount_id":
		return int(ev.Exit.Process.FileEvent.FileFields.MountID), nil
	case "exit.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.FileEvent), nil
	case "exit.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.FileEvent), nil
	case "exit.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.FileEvent), nil
	case "exit.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.FileEvent), nil
	case "exit.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exit.Process.FileEvent.FileFields)), nil
	case "exit.file.uid":
		return int(ev.Exit.Process.FileEvent.FileFields.UID), nil
	case "exit.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exit.Process.FileEvent.FileFields), nil
	case "exit.fsgid":
		return int(ev.Exit.Process.Credentials.FSGID), nil
	case "exit.fsgroup":
		return ev.Exit.Process.Credentials.FSGroup, nil
	case "exit.fsuid":
		return int(ev.Exit.Process.Credentials.FSUID), nil
	case "exit.fsuser":
		return ev.Exit.Process.Credentials.FSUser, nil
	case "exit.gid":
		return int(ev.Exit.Process.Credentials.GID), nil
	case "exit.group":
		return ev.Exit.Process.Credentials.Group, nil
	case "exit.interpreter.file.change_time":
		return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.CTime), nil
	case "exit.interpreter.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Exit.Process.LinuxBinprm.FileEvent), nil
	case "exit.interpreter.file.gid":
		return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.GID), nil
	case "exit.interpreter.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "exit.interpreter.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "exit.interpreter.file.inode":
		return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Inode), nil
	case "exit.interpreter.file.mode":
		return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Mode), nil
	case "exit.interpreter.file.modification_time":
		return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.MTime), nil
	case "exit.interpreter.file.mount_id":
		return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.MountID), nil
	case "exit.interpreter.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.LinuxBinprm.FileEvent), nil
	case "exit.interpreter.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Exit.Process.LinuxBinprm.FileEvent), nil
	case "exit.interpreter.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.LinuxBinprm.FileEvent), nil
	case "exit.interpreter.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Exit.Process.LinuxBinprm.FileEvent), nil
	case "exit.interpreter.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields)), nil
	case "exit.interpreter.file.uid":
		return int(ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.UID), nil
	case "exit.interpreter.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Exit.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "exit.is_kworker":
		return ev.Exit.Process.PIDContext.IsKworker, nil
	case "exit.is_thread":
		return ev.Exit.Process.IsThread, nil
	case "exit.pid":
		return int(ev.Exit.Process.PIDContext.Pid), nil
	case "exit.ppid":
		return int(ev.Exit.Process.PPid), nil
	case "exit.tid":
		return int(ev.Exit.Process.PIDContext.Tid), nil
	case "exit.tty_name":
		return ev.Exit.Process.TTYName, nil
	case "exit.uid":
		return int(ev.Exit.Process.Credentials.UID), nil
	case "exit.user":
		return ev.Exit.Process.Credentials.User, nil
	case "link.file.change_time":
		return int(ev.Link.Source.FileFields.CTime), nil
	case "link.file.destination.change_time":
		return int(ev.Link.Target.FileFields.CTime), nil
	case "link.file.destination.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Link.Target), nil
	case "link.file.destination.gid":
		return int(ev.Link.Target.FileFields.GID), nil
	case "link.file.destination.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Link.Target.FileFields), nil
	case "link.file.destination.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Link.Target.FileFields), nil
	case "link.file.destination.inode":
		return int(ev.Link.Target.FileFields.Inode), nil
	case "link.file.destination.mode":
		return int(ev.Link.Target.FileFields.Mode), nil
	case "link.file.destination.modification_time":
		return int(ev.Link.Target.FileFields.MTime), nil
	case "link.file.destination.mount_id":
		return int(ev.Link.Target.FileFields.MountID), nil
	case "link.file.destination.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Target), nil
	case "link.file.destination.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Target), nil
	case "link.file.destination.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Target), nil
	case "link.file.destination.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Target), nil
	case "link.file.destination.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Link.Target.FileFields)), nil
	case "link.file.destination.uid":
		return int(ev.Link.Target.FileFields.UID), nil
	case "link.file.destination.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Link.Target.FileFields), nil
	case "link.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Link.Source), nil
	case "link.file.gid":
		return int(ev.Link.Source.FileFields.GID), nil
	case "link.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Link.Source.FileFields), nil
	case "link.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Link.Source.FileFields), nil
	case "link.file.inode":
		return int(ev.Link.Source.FileFields.Inode), nil
	case "link.file.mode":
		return int(ev.Link.Source.FileFields.Mode), nil
	case "link.file.modification_time":
		return int(ev.Link.Source.FileFields.MTime), nil
	case "link.file.mount_id":
		return int(ev.Link.Source.FileFields.MountID), nil
	case "link.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Source), nil
	case "link.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Link.Source), nil
	case "link.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Source), nil
	case "link.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Link.Source), nil
	case "link.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Link.Source.FileFields)), nil
	case "link.file.uid":
		return int(ev.Link.Source.FileFields.UID), nil
	case "link.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Link.Source.FileFields), nil
	case "link.retval":
		return int(ev.Link.SyscallEvent.Retval), nil
	case "mkdir.file.change_time":
		return int(ev.Mkdir.File.FileFields.CTime), nil
	case "mkdir.file.destination.mode":
		return int(ev.Mkdir.Mode), nil
	case "mkdir.file.destination.rights":
		return int(ev.Mkdir.Mode), nil
	case "mkdir.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Mkdir.File), nil
	case "mkdir.file.gid":
		return int(ev.Mkdir.File.FileFields.GID), nil
	case "mkdir.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Mkdir.File.FileFields), nil
	case "mkdir.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Mkdir.File.FileFields), nil
	case "mkdir.file.inode":
		return int(ev.Mkdir.File.FileFields.Inode), nil
	case "mkdir.file.mode":
		return int(ev.Mkdir.File.FileFields.Mode), nil
	case "mkdir.file.modification_time":
		return int(ev.Mkdir.File.FileFields.MTime), nil
	case "mkdir.file.mount_id":
		return int(ev.Mkdir.File.FileFields.MountID), nil
	case "mkdir.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Mkdir.File), nil
	case "mkdir.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Mkdir.File), nil
	case "mkdir.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Mkdir.File), nil
	case "mkdir.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Mkdir.File), nil
	case "mkdir.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Mkdir.File.FileFields)), nil
	case "mkdir.file.uid":
		return int(ev.Mkdir.File.FileFields.UID), nil
	case "mkdir.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Mkdir.File.FileFields), nil
	case "mkdir.retval":
		return int(ev.Mkdir.SyscallEvent.Retval), nil
	case "mount.fs_type":
		return ev.Mount.Mount.FSType, nil
	case "mount.mountpoint.path":
		return ev.FieldHandlers.ResolveMountPointPath(ev, &ev.Mount), nil
	case "mount.retval":
		return int(ev.Mount.SyscallEvent.Retval), nil
	case "mount.source.path":
		return ev.FieldHandlers.ResolveMountSourcePath(ev, &ev.Mount), nil
	case "network.destination.ip":
		return ev.NetworkContext.Destination.IPNet, nil
	case "network.destination.port":
		return int(ev.NetworkContext.Destination.Port), nil
	case "network.device.ifindex":
		return int(ev.NetworkContext.Device.IfIndex), nil
	case "network.device.ifname":
		return ev.FieldHandlers.ResolveNetworkDeviceIfName(ev, &ev.NetworkContext.Device), nil
	case "network.l3_protocol":
		return int(ev.NetworkContext.L3Protocol), nil
	case "network.l4_protocol":
		return int(ev.NetworkContext.L4Protocol), nil
	case "network.size":
		return int(ev.NetworkContext.Size), nil
	case "network.source.ip":
		return ev.NetworkContext.Source.IPNet, nil
	case "network.source.port":
		return int(ev.NetworkContext.Source.Port), nil
	case "open.file.change_time":
		return int(ev.Open.File.FileFields.CTime), nil
	case "open.file.destination.mode":
		return int(ev.Open.Mode), nil
	case "open.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Open.File), nil
	case "open.file.gid":
		return int(ev.Open.File.FileFields.GID), nil
	case "open.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Open.File.FileFields), nil
	case "open.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Open.File.FileFields), nil
	case "open.file.inode":
		return int(ev.Open.File.FileFields.Inode), nil
	case "open.file.mode":
		return int(ev.Open.File.FileFields.Mode), nil
	case "open.file.modification_time":
		return int(ev.Open.File.FileFields.MTime), nil
	case "open.file.mount_id":
		return int(ev.Open.File.FileFields.MountID), nil
	case "open.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Open.File), nil
	case "open.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Open.File), nil
	case "open.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Open.File), nil
	case "open.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Open.File), nil
	case "open.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Open.File.FileFields)), nil
	case "open.file.uid":
		return int(ev.Open.File.FileFields.UID), nil
	case "open.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Open.File.FileFields), nil
	case "open.flags":
		return int(ev.Open.Flags), nil
	case "open.retval":
		return int(ev.Open.SyscallEvent.Retval), nil
	case "process.ancestors.args":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgs(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.args_flags":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgsFlags(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.args_options":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgsOptions(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.args_truncated":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.argv":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgv(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.argv0":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgv0(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.cap_effective":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.CapEffective)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.cap_permitted":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.CapPermitted)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.comm":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Comm
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.container.id":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.ContainerID
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.cookie":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Cookie)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.created_at":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &element.ProcessContext.Process))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.egid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.EGID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.egroup":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.EGroup
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.envp":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessEnvp(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.envs":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessEnvs(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.envs_truncated":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.euid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.EUID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.euser":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.EUser
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.change_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.CTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.filesystem":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.gid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.GID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.group":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.in_upper_layer":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.inode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.Inode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.mode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.Mode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.modification_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.MTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.mount_id":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.MountID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.name":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.name.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.path":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.path.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.rights":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.FileEvent.FileFields))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.uid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.UID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.file.user":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.fsgid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.FSGID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.fsgroup":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.FSGroup
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.fsuid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.FSUID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.fsuser":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.FSUser
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.gid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.GID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.group":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.Group
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.change_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.filesystem":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.gid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.group":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.in_upper_layer":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.inode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.mode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.modification_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.mount_id":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.name":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.name.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.path":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.path.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.rights":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.uid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.interpreter.file.user":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.is_kworker":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.PIDContext.IsKworker
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.is_thread":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.IsThread
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.pid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.PIDContext.Pid)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.ppid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.PPid)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.tid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.PIDContext.Tid)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.tty_name":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.TTYName
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.uid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.UID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.ancestors.user":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.User
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "process.args":
		return ev.FieldHandlers.ResolveProcessArgs(ev, &ev.ProcessContext.Process), nil
	case "process.args_flags":
		return ev.FieldHandlers.ResolveProcessArgsFlags(ev, &ev.ProcessContext.Process), nil
	case "process.args_options":
		return ev.FieldHandlers.ResolveProcessArgsOptions(ev, &ev.ProcessContext.Process), nil
	case "process.args_truncated":
		return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &ev.ProcessContext.Process), nil
	case "process.argv":
		return ev.FieldHandlers.ResolveProcessArgv(ev, &ev.ProcessContext.Process), nil
	case "process.argv0":
		return ev.FieldHandlers.ResolveProcessArgv0(ev, &ev.ProcessContext.Process), nil
	case "process.cap_effective":
		return int(ev.ProcessContext.Process.Credentials.CapEffective), nil
	case "process.cap_permitted":
		return int(ev.ProcessContext.Process.Credentials.CapPermitted), nil
	case "process.comm":
		return ev.ProcessContext.Process.Comm, nil
	case "process.container.id":
		return ev.ProcessContext.Process.ContainerID, nil
	case "process.cookie":
		return int(ev.ProcessContext.Process.Cookie), nil
	case "process.created_at":
		return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &ev.ProcessContext.Process)), nil
	case "process.egid":
		return int(ev.ProcessContext.Process.Credentials.EGID), nil
	case "process.egroup":
		return ev.ProcessContext.Process.Credentials.EGroup, nil
	case "process.envp":
		return ev.FieldHandlers.ResolveProcessEnvp(ev, &ev.ProcessContext.Process), nil
	case "process.envs":
		return ev.FieldHandlers.ResolveProcessEnvs(ev, &ev.ProcessContext.Process), nil
	case "process.envs_truncated":
		return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &ev.ProcessContext.Process), nil
	case "process.euid":
		return int(ev.ProcessContext.Process.Credentials.EUID), nil
	case "process.euser":
		return ev.ProcessContext.Process.Credentials.EUser, nil
	case "process.file.change_time":
		return int(ev.ProcessContext.Process.FileEvent.FileFields.CTime), nil
	case "process.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Process.FileEvent), nil
	case "process.file.gid":
		return int(ev.ProcessContext.Process.FileEvent.FileFields.GID), nil
	case "process.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Process.FileEvent.FileFields), nil
	case "process.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Process.FileEvent.FileFields), nil
	case "process.file.inode":
		return int(ev.ProcessContext.Process.FileEvent.FileFields.Inode), nil
	case "process.file.mode":
		return int(ev.ProcessContext.Process.FileEvent.FileFields.Mode), nil
	case "process.file.modification_time":
		return int(ev.ProcessContext.Process.FileEvent.FileFields.MTime), nil
	case "process.file.mount_id":
		return int(ev.ProcessContext.Process.FileEvent.FileFields.MountID), nil
	case "process.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.FileEvent), nil
	case "process.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.FileEvent), nil
	case "process.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.FileEvent), nil
	case "process.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.FileEvent), nil
	case "process.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Process.FileEvent.FileFields)), nil
	case "process.file.uid":
		return int(ev.ProcessContext.Process.FileEvent.FileFields.UID), nil
	case "process.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Process.FileEvent.FileFields), nil
	case "process.fsgid":
		return int(ev.ProcessContext.Process.Credentials.FSGID), nil
	case "process.fsgroup":
		return ev.ProcessContext.Process.Credentials.FSGroup, nil
	case "process.fsuid":
		return int(ev.ProcessContext.Process.Credentials.FSUID), nil
	case "process.fsuser":
		return ev.ProcessContext.Process.Credentials.FSUser, nil
	case "process.gid":
		return int(ev.ProcessContext.Process.Credentials.GID), nil
	case "process.group":
		return ev.ProcessContext.Process.Credentials.Group, nil
	case "process.interpreter.file.change_time":
		return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime), nil
	case "process.interpreter.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent), nil
	case "process.interpreter.file.gid":
		return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID), nil
	case "process.interpreter.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "process.interpreter.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "process.interpreter.file.inode":
		return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode), nil
	case "process.interpreter.file.mode":
		return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode), nil
	case "process.interpreter.file.modification_time":
		return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime), nil
	case "process.interpreter.file.mount_id":
		return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID), nil
	case "process.interpreter.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent), nil
	case "process.interpreter.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent), nil
	case "process.interpreter.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent), nil
	case "process.interpreter.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent), nil
	case "process.interpreter.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)), nil
	case "process.interpreter.file.uid":
		return int(ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID), nil
	case "process.interpreter.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "process.is_kworker":
		return ev.ProcessContext.Process.PIDContext.IsKworker, nil
	case "process.is_thread":
		return ev.ProcessContext.Process.IsThread, nil
	case "process.parent.args":
		return ev.FieldHandlers.ResolveProcessArgs(ev, ev.ProcessContext.Parent), nil
	case "process.parent.args_flags":
		return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.ProcessContext.Parent), nil
	case "process.parent.args_options":
		return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.ProcessContext.Parent), nil
	case "process.parent.args_truncated":
		return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.ProcessContext.Parent), nil
	case "process.parent.argv":
		return ev.FieldHandlers.ResolveProcessArgv(ev, ev.ProcessContext.Parent), nil
	case "process.parent.argv0":
		return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.ProcessContext.Parent), nil
	case "process.parent.cap_effective":
		return int(ev.ProcessContext.Parent.Credentials.CapEffective), nil
	case "process.parent.cap_permitted":
		return int(ev.ProcessContext.Parent.Credentials.CapPermitted), nil
	case "process.parent.comm":
		return ev.ProcessContext.Parent.Comm, nil
	case "process.parent.container.id":
		return ev.ProcessContext.Parent.ContainerID, nil
	case "process.parent.cookie":
		return int(ev.ProcessContext.Parent.Cookie), nil
	case "process.parent.created_at":
		return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.ProcessContext.Parent)), nil
	case "process.parent.egid":
		return int(ev.ProcessContext.Parent.Credentials.EGID), nil
	case "process.parent.egroup":
		return ev.ProcessContext.Parent.Credentials.EGroup, nil
	case "process.parent.envp":
		return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.ProcessContext.Parent), nil
	case "process.parent.envs":
		return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.ProcessContext.Parent), nil
	case "process.parent.envs_truncated":
		return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.ProcessContext.Parent), nil
	case "process.parent.euid":
		return int(ev.ProcessContext.Parent.Credentials.EUID), nil
	case "process.parent.euser":
		return ev.ProcessContext.Parent.Credentials.EUser, nil
	case "process.parent.file.change_time":
		return int(ev.ProcessContext.Parent.FileEvent.FileFields.CTime), nil
	case "process.parent.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Parent.FileEvent), nil
	case "process.parent.file.gid":
		return int(ev.ProcessContext.Parent.FileEvent.FileFields.GID), nil
	case "process.parent.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Parent.FileEvent.FileFields), nil
	case "process.parent.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Parent.FileEvent.FileFields), nil
	case "process.parent.file.inode":
		return int(ev.ProcessContext.Parent.FileEvent.FileFields.Inode), nil
	case "process.parent.file.mode":
		return int(ev.ProcessContext.Parent.FileEvent.FileFields.Mode), nil
	case "process.parent.file.modification_time":
		return int(ev.ProcessContext.Parent.FileEvent.FileFields.MTime), nil
	case "process.parent.file.mount_id":
		return int(ev.ProcessContext.Parent.FileEvent.FileFields.MountID), nil
	case "process.parent.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.FileEvent), nil
	case "process.parent.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.FileEvent), nil
	case "process.parent.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.FileEvent), nil
	case "process.parent.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.FileEvent), nil
	case "process.parent.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Parent.FileEvent.FileFields)), nil
	case "process.parent.file.uid":
		return int(ev.ProcessContext.Parent.FileEvent.FileFields.UID), nil
	case "process.parent.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Parent.FileEvent.FileFields), nil
	case "process.parent.fsgid":
		return int(ev.ProcessContext.Parent.Credentials.FSGID), nil
	case "process.parent.fsgroup":
		return ev.ProcessContext.Parent.Credentials.FSGroup, nil
	case "process.parent.fsuid":
		return int(ev.ProcessContext.Parent.Credentials.FSUID), nil
	case "process.parent.fsuser":
		return ev.ProcessContext.Parent.Credentials.FSUser, nil
	case "process.parent.gid":
		return int(ev.ProcessContext.Parent.Credentials.GID), nil
	case "process.parent.group":
		return ev.ProcessContext.Parent.Credentials.Group, nil
	case "process.parent.interpreter.file.change_time":
		return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.CTime), nil
	case "process.parent.interpreter.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent), nil
	case "process.parent.interpreter.file.gid":
		return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.GID), nil
	case "process.parent.interpreter.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields), nil
	case "process.parent.interpreter.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields), nil
	case "process.parent.interpreter.file.inode":
		return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Inode), nil
	case "process.parent.interpreter.file.mode":
		return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Mode), nil
	case "process.parent.interpreter.file.modification_time":
		return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MTime), nil
	case "process.parent.interpreter.file.mount_id":
		return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MountID), nil
	case "process.parent.interpreter.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent), nil
	case "process.parent.interpreter.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent), nil
	case "process.parent.interpreter.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent), nil
	case "process.parent.interpreter.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent), nil
	case "process.parent.interpreter.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields)), nil
	case "process.parent.interpreter.file.uid":
		return int(ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.UID), nil
	case "process.parent.interpreter.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields), nil
	case "process.parent.is_kworker":
		return ev.ProcessContext.Parent.PIDContext.IsKworker, nil
	case "process.parent.is_thread":
		return ev.ProcessContext.Parent.IsThread, nil
	case "process.parent.pid":
		return int(ev.ProcessContext.Parent.PIDContext.Pid), nil
	case "process.parent.ppid":
		return int(ev.ProcessContext.Parent.PPid), nil
	case "process.parent.tid":
		return int(ev.ProcessContext.Parent.PIDContext.Tid), nil
	case "process.parent.tty_name":
		return ev.ProcessContext.Parent.TTYName, nil
	case "process.parent.uid":
		return int(ev.ProcessContext.Parent.Credentials.UID), nil
	case "process.parent.user":
		return ev.ProcessContext.Parent.Credentials.User, nil
	case "process.pid":
		return int(ev.ProcessContext.Process.PIDContext.Pid), nil
	case "process.ppid":
		return int(ev.ProcessContext.Process.PPid), nil
	case "process.tid":
		return int(ev.ProcessContext.Process.PIDContext.Tid), nil
	case "process.tty_name":
		return ev.ProcessContext.Process.TTYName, nil
	case "process.uid":
		return int(ev.ProcessContext.Process.Credentials.UID), nil
	case "process.user":
		return ev.ProcessContext.Process.Credentials.User, nil
	case "removexattr.file.change_time":
		return int(ev.RemoveXAttr.File.FileFields.CTime), nil
	case "removexattr.file.destination.name":
		return ev.FieldHandlers.ResolveXAttrName(ev, &ev.RemoveXAttr), nil
	case "removexattr.file.destination.namespace":
		return ev.FieldHandlers.ResolveXAttrNamespace(ev, &ev.RemoveXAttr), nil
	case "removexattr.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.RemoveXAttr.File), nil
	case "removexattr.file.gid":
		return int(ev.RemoveXAttr.File.FileFields.GID), nil
	case "removexattr.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.RemoveXAttr.File.FileFields), nil
	case "removexattr.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.RemoveXAttr.File.FileFields), nil
	case "removexattr.file.inode":
		return int(ev.RemoveXAttr.File.FileFields.Inode), nil
	case "removexattr.file.mode":
		return int(ev.RemoveXAttr.File.FileFields.Mode), nil
	case "removexattr.file.modification_time":
		return int(ev.RemoveXAttr.File.FileFields.MTime), nil
	case "removexattr.file.mount_id":
		return int(ev.RemoveXAttr.File.FileFields.MountID), nil
	case "removexattr.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.RemoveXAttr.File), nil
	case "removexattr.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.RemoveXAttr.File), nil
	case "removexattr.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.RemoveXAttr.File), nil
	case "removexattr.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.RemoveXAttr.File), nil
	case "removexattr.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.RemoveXAttr.File.FileFields)), nil
	case "removexattr.file.uid":
		return int(ev.RemoveXAttr.File.FileFields.UID), nil
	case "removexattr.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.RemoveXAttr.File.FileFields), nil
	case "removexattr.retval":
		return int(ev.RemoveXAttr.SyscallEvent.Retval), nil
	case "rename.file.change_time":
		return int(ev.Rename.Old.FileFields.CTime), nil
	case "rename.file.destination.change_time":
		return int(ev.Rename.New.FileFields.CTime), nil
	case "rename.file.destination.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Rename.New), nil
	case "rename.file.destination.gid":
		return int(ev.Rename.New.FileFields.GID), nil
	case "rename.file.destination.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Rename.New.FileFields), nil
	case "rename.file.destination.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Rename.New.FileFields), nil
	case "rename.file.destination.inode":
		return int(ev.Rename.New.FileFields.Inode), nil
	case "rename.file.destination.mode":
		return int(ev.Rename.New.FileFields.Mode), nil
	case "rename.file.destination.modification_time":
		return int(ev.Rename.New.FileFields.MTime), nil
	case "rename.file.destination.mount_id":
		return int(ev.Rename.New.FileFields.MountID), nil
	case "rename.file.destination.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.New), nil
	case "rename.file.destination.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.New), nil
	case "rename.file.destination.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.New), nil
	case "rename.file.destination.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.New), nil
	case "rename.file.destination.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Rename.New.FileFields)), nil
	case "rename.file.destination.uid":
		return int(ev.Rename.New.FileFields.UID), nil
	case "rename.file.destination.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Rename.New.FileFields), nil
	case "rename.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Rename.Old), nil
	case "rename.file.gid":
		return int(ev.Rename.Old.FileFields.GID), nil
	case "rename.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Rename.Old.FileFields), nil
	case "rename.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Rename.Old.FileFields), nil
	case "rename.file.inode":
		return int(ev.Rename.Old.FileFields.Inode), nil
	case "rename.file.mode":
		return int(ev.Rename.Old.FileFields.Mode), nil
	case "rename.file.modification_time":
		return int(ev.Rename.Old.FileFields.MTime), nil
	case "rename.file.mount_id":
		return int(ev.Rename.Old.FileFields.MountID), nil
	case "rename.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.Old), nil
	case "rename.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rename.Old), nil
	case "rename.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.Old), nil
	case "rename.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rename.Old), nil
	case "rename.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Rename.Old.FileFields)), nil
	case "rename.file.uid":
		return int(ev.Rename.Old.FileFields.UID), nil
	case "rename.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Rename.Old.FileFields), nil
	case "rename.retval":
		return int(ev.Rename.SyscallEvent.Retval), nil
	case "rmdir.file.change_time":
		return int(ev.Rmdir.File.FileFields.CTime), nil
	case "rmdir.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Rmdir.File), nil
	case "rmdir.file.gid":
		return int(ev.Rmdir.File.FileFields.GID), nil
	case "rmdir.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Rmdir.File.FileFields), nil
	case "rmdir.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Rmdir.File.FileFields), nil
	case "rmdir.file.inode":
		return int(ev.Rmdir.File.FileFields.Inode), nil
	case "rmdir.file.mode":
		return int(ev.Rmdir.File.FileFields.Mode), nil
	case "rmdir.file.modification_time":
		return int(ev.Rmdir.File.FileFields.MTime), nil
	case "rmdir.file.mount_id":
		return int(ev.Rmdir.File.FileFields.MountID), nil
	case "rmdir.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rmdir.File), nil
	case "rmdir.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Rmdir.File), nil
	case "rmdir.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rmdir.File), nil
	case "rmdir.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Rmdir.File), nil
	case "rmdir.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Rmdir.File.FileFields)), nil
	case "rmdir.file.uid":
		return int(ev.Rmdir.File.FileFields.UID), nil
	case "rmdir.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Rmdir.File.FileFields), nil
	case "rmdir.retval":
		return int(ev.Rmdir.SyscallEvent.Retval), nil
	case "setgid.egid":
		return int(ev.SetGID.EGID), nil
	case "setgid.egroup":
		return ev.FieldHandlers.ResolveSetgidEGroup(ev, &ev.SetGID), nil
	case "setgid.fsgid":
		return int(ev.SetGID.FSGID), nil
	case "setgid.fsgroup":
		return ev.FieldHandlers.ResolveSetgidFSGroup(ev, &ev.SetGID), nil
	case "setgid.gid":
		return int(ev.SetGID.GID), nil
	case "setgid.group":
		return ev.FieldHandlers.ResolveSetgidGroup(ev, &ev.SetGID), nil
	case "setuid.euid":
		return int(ev.SetUID.EUID), nil
	case "setuid.euser":
		return ev.FieldHandlers.ResolveSetuidEUser(ev, &ev.SetUID), nil
	case "setuid.fsuid":
		return int(ev.SetUID.FSUID), nil
	case "setuid.fsuser":
		return ev.FieldHandlers.ResolveSetuidFSUser(ev, &ev.SetUID), nil
	case "setuid.uid":
		return int(ev.SetUID.UID), nil
	case "setuid.user":
		return ev.FieldHandlers.ResolveSetuidUser(ev, &ev.SetUID), nil
	case "setxattr.file.change_time":
		return int(ev.SetXAttr.File.FileFields.CTime), nil
	case "setxattr.file.destination.name":
		return ev.FieldHandlers.ResolveXAttrName(ev, &ev.SetXAttr), nil
	case "setxattr.file.destination.namespace":
		return ev.FieldHandlers.ResolveXAttrNamespace(ev, &ev.SetXAttr), nil
	case "setxattr.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.SetXAttr.File), nil
	case "setxattr.file.gid":
		return int(ev.SetXAttr.File.FileFields.GID), nil
	case "setxattr.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.SetXAttr.File.FileFields), nil
	case "setxattr.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.SetXAttr.File.FileFields), nil
	case "setxattr.file.inode":
		return int(ev.SetXAttr.File.FileFields.Inode), nil
	case "setxattr.file.mode":
		return int(ev.SetXAttr.File.FileFields.Mode), nil
	case "setxattr.file.modification_time":
		return int(ev.SetXAttr.File.FileFields.MTime), nil
	case "setxattr.file.mount_id":
		return int(ev.SetXAttr.File.FileFields.MountID), nil
	case "setxattr.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.SetXAttr.File), nil
	case "setxattr.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.SetXAttr.File), nil
	case "setxattr.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.SetXAttr.File), nil
	case "setxattr.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.SetXAttr.File), nil
	case "setxattr.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.SetXAttr.File.FileFields)), nil
	case "setxattr.file.uid":
		return int(ev.SetXAttr.File.FileFields.UID), nil
	case "setxattr.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.SetXAttr.File.FileFields), nil
	case "setxattr.retval":
		return int(ev.SetXAttr.SyscallEvent.Retval), nil
	case "signal.pid":
		return int(ev.Signal.PID), nil
	case "signal.retval":
		return int(ev.Signal.SyscallEvent.Retval), nil
	case "signal.target.ancestors.args":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgs(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.args_flags":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgsFlags(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.args_options":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgsOptions(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.args_truncated":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.argv":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgv(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.argv0":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessArgv0(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.cap_effective":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.CapEffective)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.cap_permitted":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.CapPermitted)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.comm":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Comm
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.container.id":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.ContainerID
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.cookie":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Cookie)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.created_at":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &element.ProcessContext.Process))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.egid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.EGID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.egroup":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.EGroup
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.envp":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessEnvp(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.envs":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessEnvs(ev, &element.ProcessContext.Process)
			values = append(values, result...)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.envs_truncated":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &element.ProcessContext.Process)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.euid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.EUID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.euser":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.EUser
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.change_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.CTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.filesystem":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.gid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.GID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.group":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.in_upper_layer":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.inode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.Inode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.mode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.Mode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.modification_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.MTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.mount_id":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.MountID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.name":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.name.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.path":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.path.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.rights":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.FileEvent.FileFields))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.uid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.FileEvent.FileFields.UID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.file.user":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.fsgid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.FSGID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.fsgroup":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.FSGroup
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.fsuid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.FSUID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.fsuser":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.FSUser
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.gid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.GID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.group":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.Group
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.change_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.filesystem":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFilesystem(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.gid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.group":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsGroup(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.in_upper_layer":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.inode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.mode":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.modification_time":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.mount_id":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.name":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.name.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFileBasename(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.path":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.path.length":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := len(ev.FieldHandlers.ResolveFilePath(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.rights":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(ev.FieldHandlers.ResolveRights(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields))
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.uid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.interpreter.file.user":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := ev.FieldHandlers.ResolveFileFieldsUser(ev, &element.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.is_kworker":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.PIDContext.IsKworker
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.is_thread":
		var values []bool
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.IsThread
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.pid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.PIDContext.Pid)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.ppid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.PPid)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.tid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.PIDContext.Tid)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.tty_name":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.TTYName
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.uid":
		var values []int
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := int(element.ProcessContext.Process.Credentials.UID)
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.ancestors.user":
		var values []string
		ctx := eval.NewContext(ev)
		iterator := &ProcessAncestorsIterator{}
		ptr := iterator.Front(ctx)
		for ptr != nil {
			element := (*ProcessCacheEntry)(ptr)
			result := element.ProcessContext.Process.Credentials.User
			values = append(values, result)
			ptr = iterator.Next()
		}
		return values, nil
	case "signal.target.args":
		return ev.FieldHandlers.ResolveProcessArgs(ev, &ev.Signal.Target.Process), nil
	case "signal.target.args_flags":
		return ev.FieldHandlers.ResolveProcessArgsFlags(ev, &ev.Signal.Target.Process), nil
	case "signal.target.args_options":
		return ev.FieldHandlers.ResolveProcessArgsOptions(ev, &ev.Signal.Target.Process), nil
	case "signal.target.args_truncated":
		return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, &ev.Signal.Target.Process), nil
	case "signal.target.argv":
		return ev.FieldHandlers.ResolveProcessArgv(ev, &ev.Signal.Target.Process), nil
	case "signal.target.argv0":
		return ev.FieldHandlers.ResolveProcessArgv0(ev, &ev.Signal.Target.Process), nil
	case "signal.target.cap_effective":
		return int(ev.Signal.Target.Process.Credentials.CapEffective), nil
	case "signal.target.cap_permitted":
		return int(ev.Signal.Target.Process.Credentials.CapPermitted), nil
	case "signal.target.comm":
		return ev.Signal.Target.Process.Comm, nil
	case "signal.target.container.id":
		return ev.Signal.Target.Process.ContainerID, nil
	case "signal.target.cookie":
		return int(ev.Signal.Target.Process.Cookie), nil
	case "signal.target.created_at":
		return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, &ev.Signal.Target.Process)), nil
	case "signal.target.egid":
		return int(ev.Signal.Target.Process.Credentials.EGID), nil
	case "signal.target.egroup":
		return ev.Signal.Target.Process.Credentials.EGroup, nil
	case "signal.target.envp":
		return ev.FieldHandlers.ResolveProcessEnvp(ev, &ev.Signal.Target.Process), nil
	case "signal.target.envs":
		return ev.FieldHandlers.ResolveProcessEnvs(ev, &ev.Signal.Target.Process), nil
	case "signal.target.envs_truncated":
		return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, &ev.Signal.Target.Process), nil
	case "signal.target.euid":
		return int(ev.Signal.Target.Process.Credentials.EUID), nil
	case "signal.target.euser":
		return ev.Signal.Target.Process.Credentials.EUser, nil
	case "signal.target.file.change_time":
		return int(ev.Signal.Target.Process.FileEvent.FileFields.CTime), nil
	case "signal.target.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Process.FileEvent), nil
	case "signal.target.file.gid":
		return int(ev.Signal.Target.Process.FileEvent.FileFields.GID), nil
	case "signal.target.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Process.FileEvent.FileFields), nil
	case "signal.target.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Process.FileEvent.FileFields), nil
	case "signal.target.file.inode":
		return int(ev.Signal.Target.Process.FileEvent.FileFields.Inode), nil
	case "signal.target.file.mode":
		return int(ev.Signal.Target.Process.FileEvent.FileFields.Mode), nil
	case "signal.target.file.modification_time":
		return int(ev.Signal.Target.Process.FileEvent.FileFields.MTime), nil
	case "signal.target.file.mount_id":
		return int(ev.Signal.Target.Process.FileEvent.FileFields.MountID), nil
	case "signal.target.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.FileEvent), nil
	case "signal.target.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.FileEvent), nil
	case "signal.target.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.FileEvent), nil
	case "signal.target.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.FileEvent), nil
	case "signal.target.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Process.FileEvent.FileFields)), nil
	case "signal.target.file.uid":
		return int(ev.Signal.Target.Process.FileEvent.FileFields.UID), nil
	case "signal.target.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Process.FileEvent.FileFields), nil
	case "signal.target.fsgid":
		return int(ev.Signal.Target.Process.Credentials.FSGID), nil
	case "signal.target.fsgroup":
		return ev.Signal.Target.Process.Credentials.FSGroup, nil
	case "signal.target.fsuid":
		return int(ev.Signal.Target.Process.Credentials.FSUID), nil
	case "signal.target.fsuser":
		return ev.Signal.Target.Process.Credentials.FSUser, nil
	case "signal.target.gid":
		return int(ev.Signal.Target.Process.Credentials.GID), nil
	case "signal.target.group":
		return ev.Signal.Target.Process.Credentials.Group, nil
	case "signal.target.interpreter.file.change_time":
		return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.CTime), nil
	case "signal.target.interpreter.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent), nil
	case "signal.target.interpreter.file.gid":
		return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.GID), nil
	case "signal.target.interpreter.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "signal.target.interpreter.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "signal.target.interpreter.file.inode":
		return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Inode), nil
	case "signal.target.interpreter.file.mode":
		return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Mode), nil
	case "signal.target.interpreter.file.modification_time":
		return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MTime), nil
	case "signal.target.interpreter.file.mount_id":
		return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MountID), nil
	case "signal.target.interpreter.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent), nil
	case "signal.target.interpreter.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent), nil
	case "signal.target.interpreter.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent), nil
	case "signal.target.interpreter.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent), nil
	case "signal.target.interpreter.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields)), nil
	case "signal.target.interpreter.file.uid":
		return int(ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.UID), nil
	case "signal.target.interpreter.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields), nil
	case "signal.target.is_kworker":
		return ev.Signal.Target.Process.PIDContext.IsKworker, nil
	case "signal.target.is_thread":
		return ev.Signal.Target.Process.IsThread, nil
	case "signal.target.parent.args":
		return ev.FieldHandlers.ResolveProcessArgs(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.args_flags":
		return ev.FieldHandlers.ResolveProcessArgsFlags(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.args_options":
		return ev.FieldHandlers.ResolveProcessArgsOptions(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.args_truncated":
		return ev.FieldHandlers.ResolveProcessArgsTruncated(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.argv":
		return ev.FieldHandlers.ResolveProcessArgv(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.argv0":
		return ev.FieldHandlers.ResolveProcessArgv0(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.cap_effective":
		return int(ev.Signal.Target.Parent.Credentials.CapEffective), nil
	case "signal.target.parent.cap_permitted":
		return int(ev.Signal.Target.Parent.Credentials.CapPermitted), nil
	case "signal.target.parent.comm":
		return ev.Signal.Target.Parent.Comm, nil
	case "signal.target.parent.container.id":
		return ev.Signal.Target.Parent.ContainerID, nil
	case "signal.target.parent.cookie":
		return int(ev.Signal.Target.Parent.Cookie), nil
	case "signal.target.parent.created_at":
		return int(ev.FieldHandlers.ResolveProcessCreatedAt(ev, ev.Signal.Target.Parent)), nil
	case "signal.target.parent.egid":
		return int(ev.Signal.Target.Parent.Credentials.EGID), nil
	case "signal.target.parent.egroup":
		return ev.Signal.Target.Parent.Credentials.EGroup, nil
	case "signal.target.parent.envp":
		return ev.FieldHandlers.ResolveProcessEnvp(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.envs":
		return ev.FieldHandlers.ResolveProcessEnvs(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.envs_truncated":
		return ev.FieldHandlers.ResolveProcessEnvsTruncated(ev, ev.Signal.Target.Parent), nil
	case "signal.target.parent.euid":
		return int(ev.Signal.Target.Parent.Credentials.EUID), nil
	case "signal.target.parent.euser":
		return ev.Signal.Target.Parent.Credentials.EUser, nil
	case "signal.target.parent.file.change_time":
		return int(ev.Signal.Target.Parent.FileEvent.FileFields.CTime), nil
	case "signal.target.parent.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Parent.FileEvent), nil
	case "signal.target.parent.file.gid":
		return int(ev.Signal.Target.Parent.FileEvent.FileFields.GID), nil
	case "signal.target.parent.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Parent.FileEvent.FileFields), nil
	case "signal.target.parent.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Parent.FileEvent.FileFields), nil
	case "signal.target.parent.file.inode":
		return int(ev.Signal.Target.Parent.FileEvent.FileFields.Inode), nil
	case "signal.target.parent.file.mode":
		return int(ev.Signal.Target.Parent.FileEvent.FileFields.Mode), nil
	case "signal.target.parent.file.modification_time":
		return int(ev.Signal.Target.Parent.FileEvent.FileFields.MTime), nil
	case "signal.target.parent.file.mount_id":
		return int(ev.Signal.Target.Parent.FileEvent.FileFields.MountID), nil
	case "signal.target.parent.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.FileEvent), nil
	case "signal.target.parent.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.FileEvent), nil
	case "signal.target.parent.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.FileEvent), nil
	case "signal.target.parent.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.FileEvent), nil
	case "signal.target.parent.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Parent.FileEvent.FileFields)), nil
	case "signal.target.parent.file.uid":
		return int(ev.Signal.Target.Parent.FileEvent.FileFields.UID), nil
	case "signal.target.parent.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Parent.FileEvent.FileFields), nil
	case "signal.target.parent.fsgid":
		return int(ev.Signal.Target.Parent.Credentials.FSGID), nil
	case "signal.target.parent.fsgroup":
		return ev.Signal.Target.Parent.Credentials.FSGroup, nil
	case "signal.target.parent.fsuid":
		return int(ev.Signal.Target.Parent.Credentials.FSUID), nil
	case "signal.target.parent.fsuser":
		return ev.Signal.Target.Parent.Credentials.FSUser, nil
	case "signal.target.parent.gid":
		return int(ev.Signal.Target.Parent.Credentials.GID), nil
	case "signal.target.parent.group":
		return ev.Signal.Target.Parent.Credentials.Group, nil
	case "signal.target.parent.interpreter.file.change_time":
		return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.CTime), nil
	case "signal.target.parent.interpreter.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent), nil
	case "signal.target.parent.interpreter.file.gid":
		return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.GID), nil
	case "signal.target.parent.interpreter.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields), nil
	case "signal.target.parent.interpreter.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields), nil
	case "signal.target.parent.interpreter.file.inode":
		return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Inode), nil
	case "signal.target.parent.interpreter.file.mode":
		return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Mode), nil
	case "signal.target.parent.interpreter.file.modification_time":
		return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MTime), nil
	case "signal.target.parent.interpreter.file.mount_id":
		return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MountID), nil
	case "signal.target.parent.interpreter.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent), nil
	case "signal.target.parent.interpreter.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent), nil
	case "signal.target.parent.interpreter.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent), nil
	case "signal.target.parent.interpreter.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent), nil
	case "signal.target.parent.interpreter.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields)), nil
	case "signal.target.parent.interpreter.file.uid":
		return int(ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.UID), nil
	case "signal.target.parent.interpreter.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields), nil
	case "signal.target.parent.is_kworker":
		return ev.Signal.Target.Parent.PIDContext.IsKworker, nil
	case "signal.target.parent.is_thread":
		return ev.Signal.Target.Parent.IsThread, nil
	case "signal.target.parent.pid":
		return int(ev.Signal.Target.Parent.PIDContext.Pid), nil
	case "signal.target.parent.ppid":
		return int(ev.Signal.Target.Parent.PPid), nil
	case "signal.target.parent.tid":
		return int(ev.Signal.Target.Parent.PIDContext.Tid), nil
	case "signal.target.parent.tty_name":
		return ev.Signal.Target.Parent.TTYName, nil
	case "signal.target.parent.uid":
		return int(ev.Signal.Target.Parent.Credentials.UID), nil
	case "signal.target.parent.user":
		return ev.Signal.Target.Parent.Credentials.User, nil
	case "signal.target.pid":
		return int(ev.Signal.Target.Process.PIDContext.Pid), nil
	case "signal.target.ppid":
		return int(ev.Signal.Target.Process.PPid), nil
	case "signal.target.tid":
		return int(ev.Signal.Target.Process.PIDContext.Tid), nil
	case "signal.target.tty_name":
		return ev.Signal.Target.Process.TTYName, nil
	case "signal.target.uid":
		return int(ev.Signal.Target.Process.Credentials.UID), nil
	case "signal.target.user":
		return ev.Signal.Target.Process.Credentials.User, nil
	case "signal.type":
		return int(ev.Signal.Type), nil
	case "splice.file.change_time":
		return int(ev.Splice.File.FileFields.CTime), nil
	case "splice.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Splice.File), nil
	case "splice.file.gid":
		return int(ev.Splice.File.FileFields.GID), nil
	case "splice.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Splice.File.FileFields), nil
	case "splice.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Splice.File.FileFields), nil
	case "splice.file.inode":
		return int(ev.Splice.File.FileFields.Inode), nil
	case "splice.file.mode":
		return int(ev.Splice.File.FileFields.Mode), nil
	case "splice.file.modification_time":
		return int(ev.Splice.File.FileFields.MTime), nil
	case "splice.file.mount_id":
		return int(ev.Splice.File.FileFields.MountID), nil
	case "splice.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Splice.File), nil
	case "splice.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Splice.File), nil
	case "splice.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Splice.File), nil
	case "splice.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Splice.File), nil
	case "splice.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Splice.File.FileFields)), nil
	case "splice.file.uid":
		return int(ev.Splice.File.FileFields.UID), nil
	case "splice.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Splice.File.FileFields), nil
	case "splice.pipe_entry_flag":
		return int(ev.Splice.PipeEntryFlag), nil
	case "splice.pipe_exit_flag":
		return int(ev.Splice.PipeExitFlag), nil
	case "splice.retval":
		return int(ev.Splice.SyscallEvent.Retval), nil
	case "unlink.file.change_time":
		return int(ev.Unlink.File.FileFields.CTime), nil
	case "unlink.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Unlink.File), nil
	case "unlink.file.gid":
		return int(ev.Unlink.File.FileFields.GID), nil
	case "unlink.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Unlink.File.FileFields), nil
	case "unlink.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Unlink.File.FileFields), nil
	case "unlink.file.inode":
		return int(ev.Unlink.File.FileFields.Inode), nil
	case "unlink.file.mode":
		return int(ev.Unlink.File.FileFields.Mode), nil
	case "unlink.file.modification_time":
		return int(ev.Unlink.File.FileFields.MTime), nil
	case "unlink.file.mount_id":
		return int(ev.Unlink.File.FileFields.MountID), nil
	case "unlink.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Unlink.File), nil
	case "unlink.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Unlink.File), nil
	case "unlink.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Unlink.File), nil
	case "unlink.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Unlink.File), nil
	case "unlink.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Unlink.File.FileFields)), nil
	case "unlink.file.uid":
		return int(ev.Unlink.File.FileFields.UID), nil
	case "unlink.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Unlink.File.FileFields), nil
	case "unlink.flags":
		return int(ev.Unlink.Flags), nil
	case "unlink.retval":
		return int(ev.Unlink.SyscallEvent.Retval), nil
	case "utimes.file.change_time":
		return int(ev.Utimes.File.FileFields.CTime), nil
	case "utimes.file.filesystem":
		return ev.FieldHandlers.ResolveFileFilesystem(ev, &ev.Utimes.File), nil
	case "utimes.file.gid":
		return int(ev.Utimes.File.FileFields.GID), nil
	case "utimes.file.group":
		return ev.FieldHandlers.ResolveFileFieldsGroup(ev, &ev.Utimes.File.FileFields), nil
	case "utimes.file.in_upper_layer":
		return ev.FieldHandlers.ResolveFileFieldsInUpperLayer(ev, &ev.Utimes.File.FileFields), nil
	case "utimes.file.inode":
		return int(ev.Utimes.File.FileFields.Inode), nil
	case "utimes.file.mode":
		return int(ev.Utimes.File.FileFields.Mode), nil
	case "utimes.file.modification_time":
		return int(ev.Utimes.File.FileFields.MTime), nil
	case "utimes.file.mount_id":
		return int(ev.Utimes.File.FileFields.MountID), nil
	case "utimes.file.name":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Utimes.File), nil
	case "utimes.file.name.length":
		return ev.FieldHandlers.ResolveFileBasename(ev, &ev.Utimes.File), nil
	case "utimes.file.path":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Utimes.File), nil
	case "utimes.file.path.length":
		return ev.FieldHandlers.ResolveFilePath(ev, &ev.Utimes.File), nil
	case "utimes.file.rights":
		return int(ev.FieldHandlers.ResolveRights(ev, &ev.Utimes.File.FileFields)), nil
	case "utimes.file.uid":
		return int(ev.Utimes.File.FileFields.UID), nil
	case "utimes.file.user":
		return ev.FieldHandlers.ResolveFileFieldsUser(ev, &ev.Utimes.File.FileFields), nil
	case "utimes.retval":
		return int(ev.Utimes.SyscallEvent.Retval), nil
	}
	return nil, &eval.ErrFieldNotFound{Field: field}
}
func (ev *Event) GetFieldEventType(field eval.Field) (eval.EventType, error) {
	switch field {
	case "async":
		return "*", nil
	case "bind.addr.family":
		return "bind", nil
	case "bind.addr.ip":
		return "bind", nil
	case "bind.addr.port":
		return "bind", nil
	case "bind.retval":
		return "bind", nil
	case "capset.cap_effective":
		return "capset", nil
	case "capset.cap_permitted":
		return "capset", nil
	case "chmod.file.change_time":
		return "chmod", nil
	case "chmod.file.destination.mode":
		return "chmod", nil
	case "chmod.file.destination.rights":
		return "chmod", nil
	case "chmod.file.filesystem":
		return "chmod", nil
	case "chmod.file.gid":
		return "chmod", nil
	case "chmod.file.group":
		return "chmod", nil
	case "chmod.file.in_upper_layer":
		return "chmod", nil
	case "chmod.file.inode":
		return "chmod", nil
	case "chmod.file.mode":
		return "chmod", nil
	case "chmod.file.modification_time":
		return "chmod", nil
	case "chmod.file.mount_id":
		return "chmod", nil
	case "chmod.file.name":
		return "chmod", nil
	case "chmod.file.name.length":
		return "chmod", nil
	case "chmod.file.path":
		return "chmod", nil
	case "chmod.file.path.length":
		return "chmod", nil
	case "chmod.file.rights":
		return "chmod", nil
	case "chmod.file.uid":
		return "chmod", nil
	case "chmod.file.user":
		return "chmod", nil
	case "chmod.retval":
		return "chmod", nil
	case "chown.file.change_time":
		return "chown", nil
	case "chown.file.destination.gid":
		return "chown", nil
	case "chown.file.destination.group":
		return "chown", nil
	case "chown.file.destination.uid":
		return "chown", nil
	case "chown.file.destination.user":
		return "chown", nil
	case "chown.file.filesystem":
		return "chown", nil
	case "chown.file.gid":
		return "chown", nil
	case "chown.file.group":
		return "chown", nil
	case "chown.file.in_upper_layer":
		return "chown", nil
	case "chown.file.inode":
		return "chown", nil
	case "chown.file.mode":
		return "chown", nil
	case "chown.file.modification_time":
		return "chown", nil
	case "chown.file.mount_id":
		return "chown", nil
	case "chown.file.name":
		return "chown", nil
	case "chown.file.name.length":
		return "chown", nil
	case "chown.file.path":
		return "chown", nil
	case "chown.file.path.length":
		return "chown", nil
	case "chown.file.rights":
		return "chown", nil
	case "chown.file.uid":
		return "chown", nil
	case "chown.file.user":
		return "chown", nil
	case "chown.retval":
		return "chown", nil
	case "container.id":
		return "*", nil
	case "container.tags":
		return "*", nil
	case "dns.id":
		return "dns", nil
	case "dns.question.class":
		return "dns", nil
	case "dns.question.count":
		return "dns", nil
	case "dns.question.length":
		return "dns", nil
	case "dns.question.name":
		return "dns", nil
	case "dns.question.type":
		return "dns", nil
	case "exec.args":
		return "exec", nil
	case "exec.args_flags":
		return "exec", nil
	case "exec.args_options":
		return "exec", nil
	case "exec.args_truncated":
		return "exec", nil
	case "exec.argv":
		return "exec", nil
	case "exec.argv0":
		return "exec", nil
	case "exec.cap_effective":
		return "exec", nil
	case "exec.cap_permitted":
		return "exec", nil
	case "exec.comm":
		return "exec", nil
	case "exec.container.id":
		return "exec", nil
	case "exec.cookie":
		return "exec", nil
	case "exec.created_at":
		return "exec", nil
	case "exec.egid":
		return "exec", nil
	case "exec.egroup":
		return "exec", nil
	case "exec.envp":
		return "exec", nil
	case "exec.envs":
		return "exec", nil
	case "exec.envs_truncated":
		return "exec", nil
	case "exec.euid":
		return "exec", nil
	case "exec.euser":
		return "exec", nil
	case "exec.file.change_time":
		return "exec", nil
	case "exec.file.filesystem":
		return "exec", nil
	case "exec.file.gid":
		return "exec", nil
	case "exec.file.group":
		return "exec", nil
	case "exec.file.in_upper_layer":
		return "exec", nil
	case "exec.file.inode":
		return "exec", nil
	case "exec.file.mode":
		return "exec", nil
	case "exec.file.modification_time":
		return "exec", nil
	case "exec.file.mount_id":
		return "exec", nil
	case "exec.file.name":
		return "exec", nil
	case "exec.file.name.length":
		return "exec", nil
	case "exec.file.path":
		return "exec", nil
	case "exec.file.path.length":
		return "exec", nil
	case "exec.file.rights":
		return "exec", nil
	case "exec.file.uid":
		return "exec", nil
	case "exec.file.user":
		return "exec", nil
	case "exec.fsgid":
		return "exec", nil
	case "exec.fsgroup":
		return "exec", nil
	case "exec.fsuid":
		return "exec", nil
	case "exec.fsuser":
		return "exec", nil
	case "exec.gid":
		return "exec", nil
	case "exec.group":
		return "exec", nil
	case "exec.interpreter.file.change_time":
		return "exec", nil
	case "exec.interpreter.file.filesystem":
		return "exec", nil
	case "exec.interpreter.file.gid":
		return "exec", nil
	case "exec.interpreter.file.group":
		return "exec", nil
	case "exec.interpreter.file.in_upper_layer":
		return "exec", nil
	case "exec.interpreter.file.inode":
		return "exec", nil
	case "exec.interpreter.file.mode":
		return "exec", nil
	case "exec.interpreter.file.modification_time":
		return "exec", nil
	case "exec.interpreter.file.mount_id":
		return "exec", nil
	case "exec.interpreter.file.name":
		return "exec", nil
	case "exec.interpreter.file.name.length":
		return "exec", nil
	case "exec.interpreter.file.path":
		return "exec", nil
	case "exec.interpreter.file.path.length":
		return "exec", nil
	case "exec.interpreter.file.rights":
		return "exec", nil
	case "exec.interpreter.file.uid":
		return "exec", nil
	case "exec.interpreter.file.user":
		return "exec", nil
	case "exec.is_kworker":
		return "exec", nil
	case "exec.is_thread":
		return "exec", nil
	case "exec.pid":
		return "exec", nil
	case "exec.ppid":
		return "exec", nil
	case "exec.tid":
		return "exec", nil
	case "exec.tty_name":
		return "exec", nil
	case "exec.uid":
		return "exec", nil
	case "exec.user":
		return "exec", nil
	case "exit.args":
		return "exit", nil
	case "exit.args_flags":
		return "exit", nil
	case "exit.args_options":
		return "exit", nil
	case "exit.args_truncated":
		return "exit", nil
	case "exit.argv":
		return "exit", nil
	case "exit.argv0":
		return "exit", nil
	case "exit.cap_effective":
		return "exit", nil
	case "exit.cap_permitted":
		return "exit", nil
	case "exit.cause":
		return "exit", nil
	case "exit.code":
		return "exit", nil
	case "exit.comm":
		return "exit", nil
	case "exit.container.id":
		return "exit", nil
	case "exit.cookie":
		return "exit", nil
	case "exit.created_at":
		return "exit", nil
	case "exit.egid":
		return "exit", nil
	case "exit.egroup":
		return "exit", nil
	case "exit.envp":
		return "exit", nil
	case "exit.envs":
		return "exit", nil
	case "exit.envs_truncated":
		return "exit", nil
	case "exit.euid":
		return "exit", nil
	case "exit.euser":
		return "exit", nil
	case "exit.file.change_time":
		return "exit", nil
	case "exit.file.filesystem":
		return "exit", nil
	case "exit.file.gid":
		return "exit", nil
	case "exit.file.group":
		return "exit", nil
	case "exit.file.in_upper_layer":
		return "exit", nil
	case "exit.file.inode":
		return "exit", nil
	case "exit.file.mode":
		return "exit", nil
	case "exit.file.modification_time":
		return "exit", nil
	case "exit.file.mount_id":
		return "exit", nil
	case "exit.file.name":
		return "exit", nil
	case "exit.file.name.length":
		return "exit", nil
	case "exit.file.path":
		return "exit", nil
	case "exit.file.path.length":
		return "exit", nil
	case "exit.file.rights":
		return "exit", nil
	case "exit.file.uid":
		return "exit", nil
	case "exit.file.user":
		return "exit", nil
	case "exit.fsgid":
		return "exit", nil
	case "exit.fsgroup":
		return "exit", nil
	case "exit.fsuid":
		return "exit", nil
	case "exit.fsuser":
		return "exit", nil
	case "exit.gid":
		return "exit", nil
	case "exit.group":
		return "exit", nil
	case "exit.interpreter.file.change_time":
		return "exit", nil
	case "exit.interpreter.file.filesystem":
		return "exit", nil
	case "exit.interpreter.file.gid":
		return "exit", nil
	case "exit.interpreter.file.group":
		return "exit", nil
	case "exit.interpreter.file.in_upper_layer":
		return "exit", nil
	case "exit.interpreter.file.inode":
		return "exit", nil
	case "exit.interpreter.file.mode":
		return "exit", nil
	case "exit.interpreter.file.modification_time":
		return "exit", nil
	case "exit.interpreter.file.mount_id":
		return "exit", nil
	case "exit.interpreter.file.name":
		return "exit", nil
	case "exit.interpreter.file.name.length":
		return "exit", nil
	case "exit.interpreter.file.path":
		return "exit", nil
	case "exit.interpreter.file.path.length":
		return "exit", nil
	case "exit.interpreter.file.rights":
		return "exit", nil
	case "exit.interpreter.file.uid":
		return "exit", nil
	case "exit.interpreter.file.user":
		return "exit", nil
	case "exit.is_kworker":
		return "exit", nil
	case "exit.is_thread":
		return "exit", nil
	case "exit.pid":
		return "exit", nil
	case "exit.ppid":
		return "exit", nil
	case "exit.tid":
		return "exit", nil
	case "exit.tty_name":
		return "exit", nil
	case "exit.uid":
		return "exit", nil
	case "exit.user":
		return "exit", nil
	case "link.file.change_time":
		return "link", nil
	case "link.file.destination.change_time":
		return "link", nil
	case "link.file.destination.filesystem":
		return "link", nil
	case "link.file.destination.gid":
		return "link", nil
	case "link.file.destination.group":
		return "link", nil
	case "link.file.destination.in_upper_layer":
		return "link", nil
	case "link.file.destination.inode":
		return "link", nil
	case "link.file.destination.mode":
		return "link", nil
	case "link.file.destination.modification_time":
		return "link", nil
	case "link.file.destination.mount_id":
		return "link", nil
	case "link.file.destination.name":
		return "link", nil
	case "link.file.destination.name.length":
		return "link", nil
	case "link.file.destination.path":
		return "link", nil
	case "link.file.destination.path.length":
		return "link", nil
	case "link.file.destination.rights":
		return "link", nil
	case "link.file.destination.uid":
		return "link", nil
	case "link.file.destination.user":
		return "link", nil
	case "link.file.filesystem":
		return "link", nil
	case "link.file.gid":
		return "link", nil
	case "link.file.group":
		return "link", nil
	case "link.file.in_upper_layer":
		return "link", nil
	case "link.file.inode":
		return "link", nil
	case "link.file.mode":
		return "link", nil
	case "link.file.modification_time":
		return "link", nil
	case "link.file.mount_id":
		return "link", nil
	case "link.file.name":
		return "link", nil
	case "link.file.name.length":
		return "link", nil
	case "link.file.path":
		return "link", nil
	case "link.file.path.length":
		return "link", nil
	case "link.file.rights":
		return "link", nil
	case "link.file.uid":
		return "link", nil
	case "link.file.user":
		return "link", nil
	case "link.retval":
		return "link", nil
	case "mkdir.file.change_time":
		return "mkdir", nil
	case "mkdir.file.destination.mode":
		return "mkdir", nil
	case "mkdir.file.destination.rights":
		return "mkdir", nil
	case "mkdir.file.filesystem":
		return "mkdir", nil
	case "mkdir.file.gid":
		return "mkdir", nil
	case "mkdir.file.group":
		return "mkdir", nil
	case "mkdir.file.in_upper_layer":
		return "mkdir", nil
	case "mkdir.file.inode":
		return "mkdir", nil
	case "mkdir.file.mode":
		return "mkdir", nil
	case "mkdir.file.modification_time":
		return "mkdir", nil
	case "mkdir.file.mount_id":
		return "mkdir", nil
	case "mkdir.file.name":
		return "mkdir", nil
	case "mkdir.file.name.length":
		return "mkdir", nil
	case "mkdir.file.path":
		return "mkdir", nil
	case "mkdir.file.path.length":
		return "mkdir", nil
	case "mkdir.file.rights":
		return "mkdir", nil
	case "mkdir.file.uid":
		return "mkdir", nil
	case "mkdir.file.user":
		return "mkdir", nil
	case "mkdir.retval":
		return "mkdir", nil
	case "mount.fs_type":
		return "mount", nil
	case "mount.mountpoint.path":
		return "mount", nil
	case "mount.retval":
		return "mount", nil
	case "mount.source.path":
		return "mount", nil
	case "network.destination.ip":
		return "*", nil
	case "network.destination.port":
		return "*", nil
	case "network.device.ifindex":
		return "*", nil
	case "network.device.ifname":
		return "*", nil
	case "network.l3_protocol":
		return "*", nil
	case "network.l4_protocol":
		return "*", nil
	case "network.size":
		return "*", nil
	case "network.source.ip":
		return "*", nil
	case "network.source.port":
		return "*", nil
	case "open.file.change_time":
		return "open", nil
	case "open.file.destination.mode":
		return "open", nil
	case "open.file.filesystem":
		return "open", nil
	case "open.file.gid":
		return "open", nil
	case "open.file.group":
		return "open", nil
	case "open.file.in_upper_layer":
		return "open", nil
	case "open.file.inode":
		return "open", nil
	case "open.file.mode":
		return "open", nil
	case "open.file.modification_time":
		return "open", nil
	case "open.file.mount_id":
		return "open", nil
	case "open.file.name":
		return "open", nil
	case "open.file.name.length":
		return "open", nil
	case "open.file.path":
		return "open", nil
	case "open.file.path.length":
		return "open", nil
	case "open.file.rights":
		return "open", nil
	case "open.file.uid":
		return "open", nil
	case "open.file.user":
		return "open", nil
	case "open.flags":
		return "open", nil
	case "open.retval":
		return "open", nil
	case "process.ancestors.args":
		return "*", nil
	case "process.ancestors.args_flags":
		return "*", nil
	case "process.ancestors.args_options":
		return "*", nil
	case "process.ancestors.args_truncated":
		return "*", nil
	case "process.ancestors.argv":
		return "*", nil
	case "process.ancestors.argv0":
		return "*", nil
	case "process.ancestors.cap_effective":
		return "*", nil
	case "process.ancestors.cap_permitted":
		return "*", nil
	case "process.ancestors.comm":
		return "*", nil
	case "process.ancestors.container.id":
		return "*", nil
	case "process.ancestors.cookie":
		return "*", nil
	case "process.ancestors.created_at":
		return "*", nil
	case "process.ancestors.egid":
		return "*", nil
	case "process.ancestors.egroup":
		return "*", nil
	case "process.ancestors.envp":
		return "*", nil
	case "process.ancestors.envs":
		return "*", nil
	case "process.ancestors.envs_truncated":
		return "*", nil
	case "process.ancestors.euid":
		return "*", nil
	case "process.ancestors.euser":
		return "*", nil
	case "process.ancestors.file.change_time":
		return "*", nil
	case "process.ancestors.file.filesystem":
		return "*", nil
	case "process.ancestors.file.gid":
		return "*", nil
	case "process.ancestors.file.group":
		return "*", nil
	case "process.ancestors.file.in_upper_layer":
		return "*", nil
	case "process.ancestors.file.inode":
		return "*", nil
	case "process.ancestors.file.mode":
		return "*", nil
	case "process.ancestors.file.modification_time":
		return "*", nil
	case "process.ancestors.file.mount_id":
		return "*", nil
	case "process.ancestors.file.name":
		return "*", nil
	case "process.ancestors.file.name.length":
		return "*", nil
	case "process.ancestors.file.path":
		return "*", nil
	case "process.ancestors.file.path.length":
		return "*", nil
	case "process.ancestors.file.rights":
		return "*", nil
	case "process.ancestors.file.uid":
		return "*", nil
	case "process.ancestors.file.user":
		return "*", nil
	case "process.ancestors.fsgid":
		return "*", nil
	case "process.ancestors.fsgroup":
		return "*", nil
	case "process.ancestors.fsuid":
		return "*", nil
	case "process.ancestors.fsuser":
		return "*", nil
	case "process.ancestors.gid":
		return "*", nil
	case "process.ancestors.group":
		return "*", nil
	case "process.ancestors.interpreter.file.change_time":
		return "*", nil
	case "process.ancestors.interpreter.file.filesystem":
		return "*", nil
	case "process.ancestors.interpreter.file.gid":
		return "*", nil
	case "process.ancestors.interpreter.file.group":
		return "*", nil
	case "process.ancestors.interpreter.file.in_upper_layer":
		return "*", nil
	case "process.ancestors.interpreter.file.inode":
		return "*", nil
	case "process.ancestors.interpreter.file.mode":
		return "*", nil
	case "process.ancestors.interpreter.file.modification_time":
		return "*", nil
	case "process.ancestors.interpreter.file.mount_id":
		return "*", nil
	case "process.ancestors.interpreter.file.name":
		return "*", nil
	case "process.ancestors.interpreter.file.name.length":
		return "*", nil
	case "process.ancestors.interpreter.file.path":
		return "*", nil
	case "process.ancestors.interpreter.file.path.length":
		return "*", nil
	case "process.ancestors.interpreter.file.rights":
		return "*", nil
	case "process.ancestors.interpreter.file.uid":
		return "*", nil
	case "process.ancestors.interpreter.file.user":
		return "*", nil
	case "process.ancestors.is_kworker":
		return "*", nil
	case "process.ancestors.is_thread":
		return "*", nil
	case "process.ancestors.pid":
		return "*", nil
	case "process.ancestors.ppid":
		return "*", nil
	case "process.ancestors.tid":
		return "*", nil
	case "process.ancestors.tty_name":
		return "*", nil
	case "process.ancestors.uid":
		return "*", nil
	case "process.ancestors.user":
		return "*", nil
	case "process.args":
		return "*", nil
	case "process.args_flags":
		return "*", nil
	case "process.args_options":
		return "*", nil
	case "process.args_truncated":
		return "*", nil
	case "process.argv":
		return "*", nil
	case "process.argv0":
		return "*", nil
	case "process.cap_effective":
		return "*", nil
	case "process.cap_permitted":
		return "*", nil
	case "process.comm":
		return "*", nil
	case "process.container.id":
		return "*", nil
	case "process.cookie":
		return "*", nil
	case "process.created_at":
		return "*", nil
	case "process.egid":
		return "*", nil
	case "process.egroup":
		return "*", nil
	case "process.envp":
		return "*", nil
	case "process.envs":
		return "*", nil
	case "process.envs_truncated":
		return "*", nil
	case "process.euid":
		return "*", nil
	case "process.euser":
		return "*", nil
	case "process.file.change_time":
		return "*", nil
	case "process.file.filesystem":
		return "*", nil
	case "process.file.gid":
		return "*", nil
	case "process.file.group":
		return "*", nil
	case "process.file.in_upper_layer":
		return "*", nil
	case "process.file.inode":
		return "*", nil
	case "process.file.mode":
		return "*", nil
	case "process.file.modification_time":
		return "*", nil
	case "process.file.mount_id":
		return "*", nil
	case "process.file.name":
		return "*", nil
	case "process.file.name.length":
		return "*", nil
	case "process.file.path":
		return "*", nil
	case "process.file.path.length":
		return "*", nil
	case "process.file.rights":
		return "*", nil
	case "process.file.uid":
		return "*", nil
	case "process.file.user":
		return "*", nil
	case "process.fsgid":
		return "*", nil
	case "process.fsgroup":
		return "*", nil
	case "process.fsuid":
		return "*", nil
	case "process.fsuser":
		return "*", nil
	case "process.gid":
		return "*", nil
	case "process.group":
		return "*", nil
	case "process.interpreter.file.change_time":
		return "*", nil
	case "process.interpreter.file.filesystem":
		return "*", nil
	case "process.interpreter.file.gid":
		return "*", nil
	case "process.interpreter.file.group":
		return "*", nil
	case "process.interpreter.file.in_upper_layer":
		return "*", nil
	case "process.interpreter.file.inode":
		return "*", nil
	case "process.interpreter.file.mode":
		return "*", nil
	case "process.interpreter.file.modification_time":
		return "*", nil
	case "process.interpreter.file.mount_id":
		return "*", nil
	case "process.interpreter.file.name":
		return "*", nil
	case "process.interpreter.file.name.length":
		return "*", nil
	case "process.interpreter.file.path":
		return "*", nil
	case "process.interpreter.file.path.length":
		return "*", nil
	case "process.interpreter.file.rights":
		return "*", nil
	case "process.interpreter.file.uid":
		return "*", nil
	case "process.interpreter.file.user":
		return "*", nil
	case "process.is_kworker":
		return "*", nil
	case "process.is_thread":
		return "*", nil
	case "process.parent.args":
		return "*", nil
	case "process.parent.args_flags":
		return "*", nil
	case "process.parent.args_options":
		return "*", nil
	case "process.parent.args_truncated":
		return "*", nil
	case "process.parent.argv":
		return "*", nil
	case "process.parent.argv0":
		return "*", nil
	case "process.parent.cap_effective":
		return "*", nil
	case "process.parent.cap_permitted":
		return "*", nil
	case "process.parent.comm":
		return "*", nil
	case "process.parent.container.id":
		return "*", nil
	case "process.parent.cookie":
		return "*", nil
	case "process.parent.created_at":
		return "*", nil
	case "process.parent.egid":
		return "*", nil
	case "process.parent.egroup":
		return "*", nil
	case "process.parent.envp":
		return "*", nil
	case "process.parent.envs":
		return "*", nil
	case "process.parent.envs_truncated":
		return "*", nil
	case "process.parent.euid":
		return "*", nil
	case "process.parent.euser":
		return "*", nil
	case "process.parent.file.change_time":
		return "*", nil
	case "process.parent.file.filesystem":
		return "*", nil
	case "process.parent.file.gid":
		return "*", nil
	case "process.parent.file.group":
		return "*", nil
	case "process.parent.file.in_upper_layer":
		return "*", nil
	case "process.parent.file.inode":
		return "*", nil
	case "process.parent.file.mode":
		return "*", nil
	case "process.parent.file.modification_time":
		return "*", nil
	case "process.parent.file.mount_id":
		return "*", nil
	case "process.parent.file.name":
		return "*", nil
	case "process.parent.file.name.length":
		return "*", nil
	case "process.parent.file.path":
		return "*", nil
	case "process.parent.file.path.length":
		return "*", nil
	case "process.parent.file.rights":
		return "*", nil
	case "process.parent.file.uid":
		return "*", nil
	case "process.parent.file.user":
		return "*", nil
	case "process.parent.fsgid":
		return "*", nil
	case "process.parent.fsgroup":
		return "*", nil
	case "process.parent.fsuid":
		return "*", nil
	case "process.parent.fsuser":
		return "*", nil
	case "process.parent.gid":
		return "*", nil
	case "process.parent.group":
		return "*", nil
	case "process.parent.interpreter.file.change_time":
		return "*", nil
	case "process.parent.interpreter.file.filesystem":
		return "*", nil
	case "process.parent.interpreter.file.gid":
		return "*", nil
	case "process.parent.interpreter.file.group":
		return "*", nil
	case "process.parent.interpreter.file.in_upper_layer":
		return "*", nil
	case "process.parent.interpreter.file.inode":
		return "*", nil
	case "process.parent.interpreter.file.mode":
		return "*", nil
	case "process.parent.interpreter.file.modification_time":
		return "*", nil
	case "process.parent.interpreter.file.mount_id":
		return "*", nil
	case "process.parent.interpreter.file.name":
		return "*", nil
	case "process.parent.interpreter.file.name.length":
		return "*", nil
	case "process.parent.interpreter.file.path":
		return "*", nil
	case "process.parent.interpreter.file.path.length":
		return "*", nil
	case "process.parent.interpreter.file.rights":
		return "*", nil
	case "process.parent.interpreter.file.uid":
		return "*", nil
	case "process.parent.interpreter.file.user":
		return "*", nil
	case "process.parent.is_kworker":
		return "*", nil
	case "process.parent.is_thread":
		return "*", nil
	case "process.parent.pid":
		return "*", nil
	case "process.parent.ppid":
		return "*", nil
	case "process.parent.tid":
		return "*", nil
	case "process.parent.tty_name":
		return "*", nil
	case "process.parent.uid":
		return "*", nil
	case "process.parent.user":
		return "*", nil
	case "process.pid":
		return "*", nil
	case "process.ppid":
		return "*", nil
	case "process.tid":
		return "*", nil
	case "process.tty_name":
		return "*", nil
	case "process.uid":
		return "*", nil
	case "process.user":
		return "*", nil
	case "removexattr.file.change_time":
		return "removexattr", nil
	case "removexattr.file.destination.name":
		return "removexattr", nil
	case "removexattr.file.destination.namespace":
		return "removexattr", nil
	case "removexattr.file.filesystem":
		return "removexattr", nil
	case "removexattr.file.gid":
		return "removexattr", nil
	case "removexattr.file.group":
		return "removexattr", nil
	case "removexattr.file.in_upper_layer":
		return "removexattr", nil
	case "removexattr.file.inode":
		return "removexattr", nil
	case "removexattr.file.mode":
		return "removexattr", nil
	case "removexattr.file.modification_time":
		return "removexattr", nil
	case "removexattr.file.mount_id":
		return "removexattr", nil
	case "removexattr.file.name":
		return "removexattr", nil
	case "removexattr.file.name.length":
		return "removexattr", nil
	case "removexattr.file.path":
		return "removexattr", nil
	case "removexattr.file.path.length":
		return "removexattr", nil
	case "removexattr.file.rights":
		return "removexattr", nil
	case "removexattr.file.uid":
		return "removexattr", nil
	case "removexattr.file.user":
		return "removexattr", nil
	case "removexattr.retval":
		return "removexattr", nil
	case "rename.file.change_time":
		return "rename", nil
	case "rename.file.destination.change_time":
		return "rename", nil
	case "rename.file.destination.filesystem":
		return "rename", nil
	case "rename.file.destination.gid":
		return "rename", nil
	case "rename.file.destination.group":
		return "rename", nil
	case "rename.file.destination.in_upper_layer":
		return "rename", nil
	case "rename.file.destination.inode":
		return "rename", nil
	case "rename.file.destination.mode":
		return "rename", nil
	case "rename.file.destination.modification_time":
		return "rename", nil
	case "rename.file.destination.mount_id":
		return "rename", nil
	case "rename.file.destination.name":
		return "rename", nil
	case "rename.file.destination.name.length":
		return "rename", nil
	case "rename.file.destination.path":
		return "rename", nil
	case "rename.file.destination.path.length":
		return "rename", nil
	case "rename.file.destination.rights":
		return "rename", nil
	case "rename.file.destination.uid":
		return "rename", nil
	case "rename.file.destination.user":
		return "rename", nil
	case "rename.file.filesystem":
		return "rename", nil
	case "rename.file.gid":
		return "rename", nil
	case "rename.file.group":
		return "rename", nil
	case "rename.file.in_upper_layer":
		return "rename", nil
	case "rename.file.inode":
		return "rename", nil
	case "rename.file.mode":
		return "rename", nil
	case "rename.file.modification_time":
		return "rename", nil
	case "rename.file.mount_id":
		return "rename", nil
	case "rename.file.name":
		return "rename", nil
	case "rename.file.name.length":
		return "rename", nil
	case "rename.file.path":
		return "rename", nil
	case "rename.file.path.length":
		return "rename", nil
	case "rename.file.rights":
		return "rename", nil
	case "rename.file.uid":
		return "rename", nil
	case "rename.file.user":
		return "rename", nil
	case "rename.retval":
		return "rename", nil
	case "rmdir.file.change_time":
		return "rmdir", nil
	case "rmdir.file.filesystem":
		return "rmdir", nil
	case "rmdir.file.gid":
		return "rmdir", nil
	case "rmdir.file.group":
		return "rmdir", nil
	case "rmdir.file.in_upper_layer":
		return "rmdir", nil
	case "rmdir.file.inode":
		return "rmdir", nil
	case "rmdir.file.mode":
		return "rmdir", nil
	case "rmdir.file.modification_time":
		return "rmdir", nil
	case "rmdir.file.mount_id":
		return "rmdir", nil
	case "rmdir.file.name":
		return "rmdir", nil
	case "rmdir.file.name.length":
		return "rmdir", nil
	case "rmdir.file.path":
		return "rmdir", nil
	case "rmdir.file.path.length":
		return "rmdir", nil
	case "rmdir.file.rights":
		return "rmdir", nil
	case "rmdir.file.uid":
		return "rmdir", nil
	case "rmdir.file.user":
		return "rmdir", nil
	case "rmdir.retval":
		return "rmdir", nil
	case "setgid.egid":
		return "setgid", nil
	case "setgid.egroup":
		return "setgid", nil
	case "setgid.fsgid":
		return "setgid", nil
	case "setgid.fsgroup":
		return "setgid", nil
	case "setgid.gid":
		return "setgid", nil
	case "setgid.group":
		return "setgid", nil
	case "setuid.euid":
		return "setuid", nil
	case "setuid.euser":
		return "setuid", nil
	case "setuid.fsuid":
		return "setuid", nil
	case "setuid.fsuser":
		return "setuid", nil
	case "setuid.uid":
		return "setuid", nil
	case "setuid.user":
		return "setuid", nil
	case "setxattr.file.change_time":
		return "setxattr", nil
	case "setxattr.file.destination.name":
		return "setxattr", nil
	case "setxattr.file.destination.namespace":
		return "setxattr", nil
	case "setxattr.file.filesystem":
		return "setxattr", nil
	case "setxattr.file.gid":
		return "setxattr", nil
	case "setxattr.file.group":
		return "setxattr", nil
	case "setxattr.file.in_upper_layer":
		return "setxattr", nil
	case "setxattr.file.inode":
		return "setxattr", nil
	case "setxattr.file.mode":
		return "setxattr", nil
	case "setxattr.file.modification_time":
		return "setxattr", nil
	case "setxattr.file.mount_id":
		return "setxattr", nil
	case "setxattr.file.name":
		return "setxattr", nil
	case "setxattr.file.name.length":
		return "setxattr", nil
	case "setxattr.file.path":
		return "setxattr", nil
	case "setxattr.file.path.length":
		return "setxattr", nil
	case "setxattr.file.rights":
		return "setxattr", nil
	case "setxattr.file.uid":
		return "setxattr", nil
	case "setxattr.file.user":
		return "setxattr", nil
	case "setxattr.retval":
		return "setxattr", nil
	case "signal.pid":
		return "signal", nil
	case "signal.retval":
		return "signal", nil
	case "signal.target.ancestors.args":
		return "signal", nil
	case "signal.target.ancestors.args_flags":
		return "signal", nil
	case "signal.target.ancestors.args_options":
		return "signal", nil
	case "signal.target.ancestors.args_truncated":
		return "signal", nil
	case "signal.target.ancestors.argv":
		return "signal", nil
	case "signal.target.ancestors.argv0":
		return "signal", nil
	case "signal.target.ancestors.cap_effective":
		return "signal", nil
	case "signal.target.ancestors.cap_permitted":
		return "signal", nil
	case "signal.target.ancestors.comm":
		return "signal", nil
	case "signal.target.ancestors.container.id":
		return "signal", nil
	case "signal.target.ancestors.cookie":
		return "signal", nil
	case "signal.target.ancestors.created_at":
		return "signal", nil
	case "signal.target.ancestors.egid":
		return "signal", nil
	case "signal.target.ancestors.egroup":
		return "signal", nil
	case "signal.target.ancestors.envp":
		return "signal", nil
	case "signal.target.ancestors.envs":
		return "signal", nil
	case "signal.target.ancestors.envs_truncated":
		return "signal", nil
	case "signal.target.ancestors.euid":
		return "signal", nil
	case "signal.target.ancestors.euser":
		return "signal", nil
	case "signal.target.ancestors.file.change_time":
		return "signal", nil
	case "signal.target.ancestors.file.filesystem":
		return "signal", nil
	case "signal.target.ancestors.file.gid":
		return "signal", nil
	case "signal.target.ancestors.file.group":
		return "signal", nil
	case "signal.target.ancestors.file.in_upper_layer":
		return "signal", nil
	case "signal.target.ancestors.file.inode":
		return "signal", nil
	case "signal.target.ancestors.file.mode":
		return "signal", nil
	case "signal.target.ancestors.file.modification_time":
		return "signal", nil
	case "signal.target.ancestors.file.mount_id":
		return "signal", nil
	case "signal.target.ancestors.file.name":
		return "signal", nil
	case "signal.target.ancestors.file.name.length":
		return "signal", nil
	case "signal.target.ancestors.file.path":
		return "signal", nil
	case "signal.target.ancestors.file.path.length":
		return "signal", nil
	case "signal.target.ancestors.file.rights":
		return "signal", nil
	case "signal.target.ancestors.file.uid":
		return "signal", nil
	case "signal.target.ancestors.file.user":
		return "signal", nil
	case "signal.target.ancestors.fsgid":
		return "signal", nil
	case "signal.target.ancestors.fsgroup":
		return "signal", nil
	case "signal.target.ancestors.fsuid":
		return "signal", nil
	case "signal.target.ancestors.fsuser":
		return "signal", nil
	case "signal.target.ancestors.gid":
		return "signal", nil
	case "signal.target.ancestors.group":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.change_time":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.filesystem":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.gid":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.group":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.in_upper_layer":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.inode":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.mode":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.modification_time":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.mount_id":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.name":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.name.length":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.path":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.path.length":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.rights":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.uid":
		return "signal", nil
	case "signal.target.ancestors.interpreter.file.user":
		return "signal", nil
	case "signal.target.ancestors.is_kworker":
		return "signal", nil
	case "signal.target.ancestors.is_thread":
		return "signal", nil
	case "signal.target.ancestors.pid":
		return "signal", nil
	case "signal.target.ancestors.ppid":
		return "signal", nil
	case "signal.target.ancestors.tid":
		return "signal", nil
	case "signal.target.ancestors.tty_name":
		return "signal", nil
	case "signal.target.ancestors.uid":
		return "signal", nil
	case "signal.target.ancestors.user":
		return "signal", nil
	case "signal.target.args":
		return "signal", nil
	case "signal.target.args_flags":
		return "signal", nil
	case "signal.target.args_options":
		return "signal", nil
	case "signal.target.args_truncated":
		return "signal", nil
	case "signal.target.argv":
		return "signal", nil
	case "signal.target.argv0":
		return "signal", nil
	case "signal.target.cap_effective":
		return "signal", nil
	case "signal.target.cap_permitted":
		return "signal", nil
	case "signal.target.comm":
		return "signal", nil
	case "signal.target.container.id":
		return "signal", nil
	case "signal.target.cookie":
		return "signal", nil
	case "signal.target.created_at":
		return "signal", nil
	case "signal.target.egid":
		return "signal", nil
	case "signal.target.egroup":
		return "signal", nil
	case "signal.target.envp":
		return "signal", nil
	case "signal.target.envs":
		return "signal", nil
	case "signal.target.envs_truncated":
		return "signal", nil
	case "signal.target.euid":
		return "signal", nil
	case "signal.target.euser":
		return "signal", nil
	case "signal.target.file.change_time":
		return "signal", nil
	case "signal.target.file.filesystem":
		return "signal", nil
	case "signal.target.file.gid":
		return "signal", nil
	case "signal.target.file.group":
		return "signal", nil
	case "signal.target.file.in_upper_layer":
		return "signal", nil
	case "signal.target.file.inode":
		return "signal", nil
	case "signal.target.file.mode":
		return "signal", nil
	case "signal.target.file.modification_time":
		return "signal", nil
	case "signal.target.file.mount_id":
		return "signal", nil
	case "signal.target.file.name":
		return "signal", nil
	case "signal.target.file.name.length":
		return "signal", nil
	case "signal.target.file.path":
		return "signal", nil
	case "signal.target.file.path.length":
		return "signal", nil
	case "signal.target.file.rights":
		return "signal", nil
	case "signal.target.file.uid":
		return "signal", nil
	case "signal.target.file.user":
		return "signal", nil
	case "signal.target.fsgid":
		return "signal", nil
	case "signal.target.fsgroup":
		return "signal", nil
	case "signal.target.fsuid":
		return "signal", nil
	case "signal.target.fsuser":
		return "signal", nil
	case "signal.target.gid":
		return "signal", nil
	case "signal.target.group":
		return "signal", nil
	case "signal.target.interpreter.file.change_time":
		return "signal", nil
	case "signal.target.interpreter.file.filesystem":
		return "signal", nil
	case "signal.target.interpreter.file.gid":
		return "signal", nil
	case "signal.target.interpreter.file.group":
		return "signal", nil
	case "signal.target.interpreter.file.in_upper_layer":
		return "signal", nil
	case "signal.target.interpreter.file.inode":
		return "signal", nil
	case "signal.target.interpreter.file.mode":
		return "signal", nil
	case "signal.target.interpreter.file.modification_time":
		return "signal", nil
	case "signal.target.interpreter.file.mount_id":
		return "signal", nil
	case "signal.target.interpreter.file.name":
		return "signal", nil
	case "signal.target.interpreter.file.name.length":
		return "signal", nil
	case "signal.target.interpreter.file.path":
		return "signal", nil
	case "signal.target.interpreter.file.path.length":
		return "signal", nil
	case "signal.target.interpreter.file.rights":
		return "signal", nil
	case "signal.target.interpreter.file.uid":
		return "signal", nil
	case "signal.target.interpreter.file.user":
		return "signal", nil
	case "signal.target.is_kworker":
		return "signal", nil
	case "signal.target.is_thread":
		return "signal", nil
	case "signal.target.parent.args":
		return "signal", nil
	case "signal.target.parent.args_flags":
		return "signal", nil
	case "signal.target.parent.args_options":
		return "signal", nil
	case "signal.target.parent.args_truncated":
		return "signal", nil
	case "signal.target.parent.argv":
		return "signal", nil
	case "signal.target.parent.argv0":
		return "signal", nil
	case "signal.target.parent.cap_effective":
		return "signal", nil
	case "signal.target.parent.cap_permitted":
		return "signal", nil
	case "signal.target.parent.comm":
		return "signal", nil
	case "signal.target.parent.container.id":
		return "signal", nil
	case "signal.target.parent.cookie":
		return "signal", nil
	case "signal.target.parent.created_at":
		return "signal", nil
	case "signal.target.parent.egid":
		return "signal", nil
	case "signal.target.parent.egroup":
		return "signal", nil
	case "signal.target.parent.envp":
		return "signal", nil
	case "signal.target.parent.envs":
		return "signal", nil
	case "signal.target.parent.envs_truncated":
		return "signal", nil
	case "signal.target.parent.euid":
		return "signal", nil
	case "signal.target.parent.euser":
		return "signal", nil
	case "signal.target.parent.file.change_time":
		return "signal", nil
	case "signal.target.parent.file.filesystem":
		return "signal", nil
	case "signal.target.parent.file.gid":
		return "signal", nil
	case "signal.target.parent.file.group":
		return "signal", nil
	case "signal.target.parent.file.in_upper_layer":
		return "signal", nil
	case "signal.target.parent.file.inode":
		return "signal", nil
	case "signal.target.parent.file.mode":
		return "signal", nil
	case "signal.target.parent.file.modification_time":
		return "signal", nil
	case "signal.target.parent.file.mount_id":
		return "signal", nil
	case "signal.target.parent.file.name":
		return "signal", nil
	case "signal.target.parent.file.name.length":
		return "signal", nil
	case "signal.target.parent.file.path":
		return "signal", nil
	case "signal.target.parent.file.path.length":
		return "signal", nil
	case "signal.target.parent.file.rights":
		return "signal", nil
	case "signal.target.parent.file.uid":
		return "signal", nil
	case "signal.target.parent.file.user":
		return "signal", nil
	case "signal.target.parent.fsgid":
		return "signal", nil
	case "signal.target.parent.fsgroup":
		return "signal", nil
	case "signal.target.parent.fsuid":
		return "signal", nil
	case "signal.target.parent.fsuser":
		return "signal", nil
	case "signal.target.parent.gid":
		return "signal", nil
	case "signal.target.parent.group":
		return "signal", nil
	case "signal.target.parent.interpreter.file.change_time":
		return "signal", nil
	case "signal.target.parent.interpreter.file.filesystem":
		return "signal", nil
	case "signal.target.parent.interpreter.file.gid":
		return "signal", nil
	case "signal.target.parent.interpreter.file.group":
		return "signal", nil
	case "signal.target.parent.interpreter.file.in_upper_layer":
		return "signal", nil
	case "signal.target.parent.interpreter.file.inode":
		return "signal", nil
	case "signal.target.parent.interpreter.file.mode":
		return "signal", nil
	case "signal.target.parent.interpreter.file.modification_time":
		return "signal", nil
	case "signal.target.parent.interpreter.file.mount_id":
		return "signal", nil
	case "signal.target.parent.interpreter.file.name":
		return "signal", nil
	case "signal.target.parent.interpreter.file.name.length":
		return "signal", nil
	case "signal.target.parent.interpreter.file.path":
		return "signal", nil
	case "signal.target.parent.interpreter.file.path.length":
		return "signal", nil
	case "signal.target.parent.interpreter.file.rights":
		return "signal", nil
	case "signal.target.parent.interpreter.file.uid":
		return "signal", nil
	case "signal.target.parent.interpreter.file.user":
		return "signal", nil
	case "signal.target.parent.is_kworker":
		return "signal", nil
	case "signal.target.parent.is_thread":
		return "signal", nil
	case "signal.target.parent.pid":
		return "signal", nil
	case "signal.target.parent.ppid":
		return "signal", nil
	case "signal.target.parent.tid":
		return "signal", nil
	case "signal.target.parent.tty_name":
		return "signal", nil
	case "signal.target.parent.uid":
		return "signal", nil
	case "signal.target.parent.user":
		return "signal", nil
	case "signal.target.pid":
		return "signal", nil
	case "signal.target.ppid":
		return "signal", nil
	case "signal.target.tid":
		return "signal", nil
	case "signal.target.tty_name":
		return "signal", nil
	case "signal.target.uid":
		return "signal", nil
	case "signal.target.user":
		return "signal", nil
	case "signal.type":
		return "signal", nil
	case "splice.file.change_time":
		return "splice", nil
	case "splice.file.filesystem":
		return "splice", nil
	case "splice.file.gid":
		return "splice", nil
	case "splice.file.group":
		return "splice", nil
	case "splice.file.in_upper_layer":
		return "splice", nil
	case "splice.file.inode":
		return "splice", nil
	case "splice.file.mode":
		return "splice", nil
	case "splice.file.modification_time":
		return "splice", nil
	case "splice.file.mount_id":
		return "splice", nil
	case "splice.file.name":
		return "splice", nil
	case "splice.file.name.length":
		return "splice", nil
	case "splice.file.path":
		return "splice", nil
	case "splice.file.path.length":
		return "splice", nil
	case "splice.file.rights":
		return "splice", nil
	case "splice.file.uid":
		return "splice", nil
	case "splice.file.user":
		return "splice", nil
	case "splice.pipe_entry_flag":
		return "splice", nil
	case "splice.pipe_exit_flag":
		return "splice", nil
	case "splice.retval":
		return "splice", nil
	case "unlink.file.change_time":
		return "unlink", nil
	case "unlink.file.filesystem":
		return "unlink", nil
	case "unlink.file.gid":
		return "unlink", nil
	case "unlink.file.group":
		return "unlink", nil
	case "unlink.file.in_upper_layer":
		return "unlink", nil
	case "unlink.file.inode":
		return "unlink", nil
	case "unlink.file.mode":
		return "unlink", nil
	case "unlink.file.modification_time":
		return "unlink", nil
	case "unlink.file.mount_id":
		return "unlink", nil
	case "unlink.file.name":
		return "unlink", nil
	case "unlink.file.name.length":
		return "unlink", nil
	case "unlink.file.path":
		return "unlink", nil
	case "unlink.file.path.length":
		return "unlink", nil
	case "unlink.file.rights":
		return "unlink", nil
	case "unlink.file.uid":
		return "unlink", nil
	case "unlink.file.user":
		return "unlink", nil
	case "unlink.flags":
		return "unlink", nil
	case "unlink.retval":
		return "unlink", nil
	case "utimes.file.change_time":
		return "utimes", nil
	case "utimes.file.filesystem":
		return "utimes", nil
	case "utimes.file.gid":
		return "utimes", nil
	case "utimes.file.group":
		return "utimes", nil
	case "utimes.file.in_upper_layer":
		return "utimes", nil
	case "utimes.file.inode":
		return "utimes", nil
	case "utimes.file.mode":
		return "utimes", nil
	case "utimes.file.modification_time":
		return "utimes", nil
	case "utimes.file.mount_id":
		return "utimes", nil
	case "utimes.file.name":
		return "utimes", nil
	case "utimes.file.name.length":
		return "utimes", nil
	case "utimes.file.path":
		return "utimes", nil
	case "utimes.file.path.length":
		return "utimes", nil
	case "utimes.file.rights":
		return "utimes", nil
	case "utimes.file.uid":
		return "utimes", nil
	case "utimes.file.user":
		return "utimes", nil
	case "utimes.retval":
		return "utimes", nil
	}
	return "", &eval.ErrFieldNotFound{Field: field}
}
func (ev *Event) GetFieldType(field eval.Field) (reflect.Kind, error) {
	switch field {
	case "async":
		return reflect.Bool, nil
	case "bind.addr.family":
		return reflect.Int, nil
	case "bind.addr.ip":
		return reflect.Struct, nil
	case "bind.addr.port":
		return reflect.Int, nil
	case "bind.retval":
		return reflect.Int, nil
	case "capset.cap_effective":
		return reflect.Int, nil
	case "capset.cap_permitted":
		return reflect.Int, nil
	case "chmod.file.change_time":
		return reflect.Int, nil
	case "chmod.file.destination.mode":
		return reflect.Int, nil
	case "chmod.file.destination.rights":
		return reflect.Int, nil
	case "chmod.file.filesystem":
		return reflect.String, nil
	case "chmod.file.gid":
		return reflect.Int, nil
	case "chmod.file.group":
		return reflect.String, nil
	case "chmod.file.in_upper_layer":
		return reflect.Bool, nil
	case "chmod.file.inode":
		return reflect.Int, nil
	case "chmod.file.mode":
		return reflect.Int, nil
	case "chmod.file.modification_time":
		return reflect.Int, nil
	case "chmod.file.mount_id":
		return reflect.Int, nil
	case "chmod.file.name":
		return reflect.String, nil
	case "chmod.file.name.length":
		return reflect.Int, nil
	case "chmod.file.path":
		return reflect.String, nil
	case "chmod.file.path.length":
		return reflect.Int, nil
	case "chmod.file.rights":
		return reflect.Int, nil
	case "chmod.file.uid":
		return reflect.Int, nil
	case "chmod.file.user":
		return reflect.String, nil
	case "chmod.retval":
		return reflect.Int, nil
	case "chown.file.change_time":
		return reflect.Int, nil
	case "chown.file.destination.gid":
		return reflect.Int, nil
	case "chown.file.destination.group":
		return reflect.String, nil
	case "chown.file.destination.uid":
		return reflect.Int, nil
	case "chown.file.destination.user":
		return reflect.String, nil
	case "chown.file.filesystem":
		return reflect.String, nil
	case "chown.file.gid":
		return reflect.Int, nil
	case "chown.file.group":
		return reflect.String, nil
	case "chown.file.in_upper_layer":
		return reflect.Bool, nil
	case "chown.file.inode":
		return reflect.Int, nil
	case "chown.file.mode":
		return reflect.Int, nil
	case "chown.file.modification_time":
		return reflect.Int, nil
	case "chown.file.mount_id":
		return reflect.Int, nil
	case "chown.file.name":
		return reflect.String, nil
	case "chown.file.name.length":
		return reflect.Int, nil
	case "chown.file.path":
		return reflect.String, nil
	case "chown.file.path.length":
		return reflect.Int, nil
	case "chown.file.rights":
		return reflect.Int, nil
	case "chown.file.uid":
		return reflect.Int, nil
	case "chown.file.user":
		return reflect.String, nil
	case "chown.retval":
		return reflect.Int, nil
	case "container.id":
		return reflect.String, nil
	case "container.tags":
		return reflect.String, nil
	case "dns.id":
		return reflect.Int, nil
	case "dns.question.class":
		return reflect.Int, nil
	case "dns.question.count":
		return reflect.Int, nil
	case "dns.question.length":
		return reflect.Int, nil
	case "dns.question.name":
		return reflect.Int, nil
	case "dns.question.type":
		return reflect.Int, nil
	case "exec.args":
		return reflect.String, nil
	case "exec.args_flags":
		return reflect.String, nil
	case "exec.args_options":
		return reflect.String, nil
	case "exec.args_truncated":
		return reflect.Bool, nil
	case "exec.argv":
		return reflect.String, nil
	case "exec.argv0":
		return reflect.String, nil
	case "exec.cap_effective":
		return reflect.Int, nil
	case "exec.cap_permitted":
		return reflect.Int, nil
	case "exec.comm":
		return reflect.String, nil
	case "exec.container.id":
		return reflect.String, nil
	case "exec.cookie":
		return reflect.Int, nil
	case "exec.created_at":
		return reflect.Int, nil
	case "exec.egid":
		return reflect.Int, nil
	case "exec.egroup":
		return reflect.String, nil
	case "exec.envp":
		return reflect.String, nil
	case "exec.envs":
		return reflect.String, nil
	case "exec.envs_truncated":
		return reflect.Bool, nil
	case "exec.euid":
		return reflect.Int, nil
	case "exec.euser":
		return reflect.String, nil
	case "exec.file.change_time":
		return reflect.Int, nil
	case "exec.file.filesystem":
		return reflect.String, nil
	case "exec.file.gid":
		return reflect.Int, nil
	case "exec.file.group":
		return reflect.String, nil
	case "exec.file.in_upper_layer":
		return reflect.Bool, nil
	case "exec.file.inode":
		return reflect.Int, nil
	case "exec.file.mode":
		return reflect.Int, nil
	case "exec.file.modification_time":
		return reflect.Int, nil
	case "exec.file.mount_id":
		return reflect.Int, nil
	case "exec.file.name":
		return reflect.String, nil
	case "exec.file.name.length":
		return reflect.Int, nil
	case "exec.file.path":
		return reflect.String, nil
	case "exec.file.path.length":
		return reflect.Int, nil
	case "exec.file.rights":
		return reflect.Int, nil
	case "exec.file.uid":
		return reflect.Int, nil
	case "exec.file.user":
		return reflect.String, nil
	case "exec.fsgid":
		return reflect.Int, nil
	case "exec.fsgroup":
		return reflect.String, nil
	case "exec.fsuid":
		return reflect.Int, nil
	case "exec.fsuser":
		return reflect.String, nil
	case "exec.gid":
		return reflect.Int, nil
	case "exec.group":
		return reflect.String, nil
	case "exec.interpreter.file.change_time":
		return reflect.Int, nil
	case "exec.interpreter.file.filesystem":
		return reflect.String, nil
	case "exec.interpreter.file.gid":
		return reflect.Int, nil
	case "exec.interpreter.file.group":
		return reflect.String, nil
	case "exec.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "exec.interpreter.file.inode":
		return reflect.Int, nil
	case "exec.interpreter.file.mode":
		return reflect.Int, nil
	case "exec.interpreter.file.modification_time":
		return reflect.Int, nil
	case "exec.interpreter.file.mount_id":
		return reflect.Int, nil
	case "exec.interpreter.file.name":
		return reflect.String, nil
	case "exec.interpreter.file.name.length":
		return reflect.Int, nil
	case "exec.interpreter.file.path":
		return reflect.String, nil
	case "exec.interpreter.file.path.length":
		return reflect.Int, nil
	case "exec.interpreter.file.rights":
		return reflect.Int, nil
	case "exec.interpreter.file.uid":
		return reflect.Int, nil
	case "exec.interpreter.file.user":
		return reflect.String, nil
	case "exec.is_kworker":
		return reflect.Bool, nil
	case "exec.is_thread":
		return reflect.Bool, nil
	case "exec.pid":
		return reflect.Int, nil
	case "exec.ppid":
		return reflect.Int, nil
	case "exec.tid":
		return reflect.Int, nil
	case "exec.tty_name":
		return reflect.String, nil
	case "exec.uid":
		return reflect.Int, nil
	case "exec.user":
		return reflect.String, nil
	case "exit.args":
		return reflect.String, nil
	case "exit.args_flags":
		return reflect.String, nil
	case "exit.args_options":
		return reflect.String, nil
	case "exit.args_truncated":
		return reflect.Bool, nil
	case "exit.argv":
		return reflect.String, nil
	case "exit.argv0":
		return reflect.String, nil
	case "exit.cap_effective":
		return reflect.Int, nil
	case "exit.cap_permitted":
		return reflect.Int, nil
	case "exit.cause":
		return reflect.Int, nil
	case "exit.code":
		return reflect.Int, nil
	case "exit.comm":
		return reflect.String, nil
	case "exit.container.id":
		return reflect.String, nil
	case "exit.cookie":
		return reflect.Int, nil
	case "exit.created_at":
		return reflect.Int, nil
	case "exit.egid":
		return reflect.Int, nil
	case "exit.egroup":
		return reflect.String, nil
	case "exit.envp":
		return reflect.String, nil
	case "exit.envs":
		return reflect.String, nil
	case "exit.envs_truncated":
		return reflect.Bool, nil
	case "exit.euid":
		return reflect.Int, nil
	case "exit.euser":
		return reflect.String, nil
	case "exit.file.change_time":
		return reflect.Int, nil
	case "exit.file.filesystem":
		return reflect.String, nil
	case "exit.file.gid":
		return reflect.Int, nil
	case "exit.file.group":
		return reflect.String, nil
	case "exit.file.in_upper_layer":
		return reflect.Bool, nil
	case "exit.file.inode":
		return reflect.Int, nil
	case "exit.file.mode":
		return reflect.Int, nil
	case "exit.file.modification_time":
		return reflect.Int, nil
	case "exit.file.mount_id":
		return reflect.Int, nil
	case "exit.file.name":
		return reflect.String, nil
	case "exit.file.name.length":
		return reflect.Int, nil
	case "exit.file.path":
		return reflect.String, nil
	case "exit.file.path.length":
		return reflect.Int, nil
	case "exit.file.rights":
		return reflect.Int, nil
	case "exit.file.uid":
		return reflect.Int, nil
	case "exit.file.user":
		return reflect.String, nil
	case "exit.fsgid":
		return reflect.Int, nil
	case "exit.fsgroup":
		return reflect.String, nil
	case "exit.fsuid":
		return reflect.Int, nil
	case "exit.fsuser":
		return reflect.String, nil
	case "exit.gid":
		return reflect.Int, nil
	case "exit.group":
		return reflect.String, nil
	case "exit.interpreter.file.change_time":
		return reflect.Int, nil
	case "exit.interpreter.file.filesystem":
		return reflect.String, nil
	case "exit.interpreter.file.gid":
		return reflect.Int, nil
	case "exit.interpreter.file.group":
		return reflect.String, nil
	case "exit.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "exit.interpreter.file.inode":
		return reflect.Int, nil
	case "exit.interpreter.file.mode":
		return reflect.Int, nil
	case "exit.interpreter.file.modification_time":
		return reflect.Int, nil
	case "exit.interpreter.file.mount_id":
		return reflect.Int, nil
	case "exit.interpreter.file.name":
		return reflect.String, nil
	case "exit.interpreter.file.name.length":
		return reflect.Int, nil
	case "exit.interpreter.file.path":
		return reflect.String, nil
	case "exit.interpreter.file.path.length":
		return reflect.Int, nil
	case "exit.interpreter.file.rights":
		return reflect.Int, nil
	case "exit.interpreter.file.uid":
		return reflect.Int, nil
	case "exit.interpreter.file.user":
		return reflect.String, nil
	case "exit.is_kworker":
		return reflect.Bool, nil
	case "exit.is_thread":
		return reflect.Bool, nil
	case "exit.pid":
		return reflect.Int, nil
	case "exit.ppid":
		return reflect.Int, nil
	case "exit.tid":
		return reflect.Int, nil
	case "exit.tty_name":
		return reflect.String, nil
	case "exit.uid":
		return reflect.Int, nil
	case "exit.user":
		return reflect.String, nil
	case "link.file.change_time":
		return reflect.Int, nil
	case "link.file.destination.change_time":
		return reflect.Int, nil
	case "link.file.destination.filesystem":
		return reflect.String, nil
	case "link.file.destination.gid":
		return reflect.Int, nil
	case "link.file.destination.group":
		return reflect.String, nil
	case "link.file.destination.in_upper_layer":
		return reflect.Bool, nil
	case "link.file.destination.inode":
		return reflect.Int, nil
	case "link.file.destination.mode":
		return reflect.Int, nil
	case "link.file.destination.modification_time":
		return reflect.Int, nil
	case "link.file.destination.mount_id":
		return reflect.Int, nil
	case "link.file.destination.name":
		return reflect.String, nil
	case "link.file.destination.name.length":
		return reflect.Int, nil
	case "link.file.destination.path":
		return reflect.String, nil
	case "link.file.destination.path.length":
		return reflect.Int, nil
	case "link.file.destination.rights":
		return reflect.Int, nil
	case "link.file.destination.uid":
		return reflect.Int, nil
	case "link.file.destination.user":
		return reflect.String, nil
	case "link.file.filesystem":
		return reflect.String, nil
	case "link.file.gid":
		return reflect.Int, nil
	case "link.file.group":
		return reflect.String, nil
	case "link.file.in_upper_layer":
		return reflect.Bool, nil
	case "link.file.inode":
		return reflect.Int, nil
	case "link.file.mode":
		return reflect.Int, nil
	case "link.file.modification_time":
		return reflect.Int, nil
	case "link.file.mount_id":
		return reflect.Int, nil
	case "link.file.name":
		return reflect.String, nil
	case "link.file.name.length":
		return reflect.Int, nil
	case "link.file.path":
		return reflect.String, nil
	case "link.file.path.length":
		return reflect.Int, nil
	case "link.file.rights":
		return reflect.Int, nil
	case "link.file.uid":
		return reflect.Int, nil
	case "link.file.user":
		return reflect.String, nil
	case "link.retval":
		return reflect.Int, nil
	case "mkdir.file.change_time":
		return reflect.Int, nil
	case "mkdir.file.destination.mode":
		return reflect.Int, nil
	case "mkdir.file.destination.rights":
		return reflect.Int, nil
	case "mkdir.file.filesystem":
		return reflect.String, nil
	case "mkdir.file.gid":
		return reflect.Int, nil
	case "mkdir.file.group":
		return reflect.String, nil
	case "mkdir.file.in_upper_layer":
		return reflect.Bool, nil
	case "mkdir.file.inode":
		return reflect.Int, nil
	case "mkdir.file.mode":
		return reflect.Int, nil
	case "mkdir.file.modification_time":
		return reflect.Int, nil
	case "mkdir.file.mount_id":
		return reflect.Int, nil
	case "mkdir.file.name":
		return reflect.String, nil
	case "mkdir.file.name.length":
		return reflect.Int, nil
	case "mkdir.file.path":
		return reflect.String, nil
	case "mkdir.file.path.length":
		return reflect.Int, nil
	case "mkdir.file.rights":
		return reflect.Int, nil
	case "mkdir.file.uid":
		return reflect.Int, nil
	case "mkdir.file.user":
		return reflect.String, nil
	case "mkdir.retval":
		return reflect.Int, nil
	case "mount.fs_type":
		return reflect.String, nil
	case "mount.mountpoint.path":
		return reflect.String, nil
	case "mount.retval":
		return reflect.Int, nil
	case "mount.source.path":
		return reflect.String, nil
	case "network.destination.ip":
		return reflect.Struct, nil
	case "network.destination.port":
		return reflect.Int, nil
	case "network.device.ifindex":
		return reflect.Int, nil
	case "network.device.ifname":
		return reflect.String, nil
	case "network.l3_protocol":
		return reflect.Int, nil
	case "network.l4_protocol":
		return reflect.Int, nil
	case "network.size":
		return reflect.Int, nil
	case "network.source.ip":
		return reflect.Struct, nil
	case "network.source.port":
		return reflect.Int, nil
	case "open.file.change_time":
		return reflect.Int, nil
	case "open.file.destination.mode":
		return reflect.Int, nil
	case "open.file.filesystem":
		return reflect.String, nil
	case "open.file.gid":
		return reflect.Int, nil
	case "open.file.group":
		return reflect.String, nil
	case "open.file.in_upper_layer":
		return reflect.Bool, nil
	case "open.file.inode":
		return reflect.Int, nil
	case "open.file.mode":
		return reflect.Int, nil
	case "open.file.modification_time":
		return reflect.Int, nil
	case "open.file.mount_id":
		return reflect.Int, nil
	case "open.file.name":
		return reflect.String, nil
	case "open.file.name.length":
		return reflect.Int, nil
	case "open.file.path":
		return reflect.String, nil
	case "open.file.path.length":
		return reflect.Int, nil
	case "open.file.rights":
		return reflect.Int, nil
	case "open.file.uid":
		return reflect.Int, nil
	case "open.file.user":
		return reflect.String, nil
	case "open.flags":
		return reflect.Int, nil
	case "open.retval":
		return reflect.Int, nil
	case "process.ancestors.args":
		return reflect.String, nil
	case "process.ancestors.args_flags":
		return reflect.String, nil
	case "process.ancestors.args_options":
		return reflect.String, nil
	case "process.ancestors.args_truncated":
		return reflect.Bool, nil
	case "process.ancestors.argv":
		return reflect.String, nil
	case "process.ancestors.argv0":
		return reflect.String, nil
	case "process.ancestors.cap_effective":
		return reflect.Int, nil
	case "process.ancestors.cap_permitted":
		return reflect.Int, nil
	case "process.ancestors.comm":
		return reflect.String, nil
	case "process.ancestors.container.id":
		return reflect.String, nil
	case "process.ancestors.cookie":
		return reflect.Int, nil
	case "process.ancestors.created_at":
		return reflect.Int, nil
	case "process.ancestors.egid":
		return reflect.Int, nil
	case "process.ancestors.egroup":
		return reflect.String, nil
	case "process.ancestors.envp":
		return reflect.String, nil
	case "process.ancestors.envs":
		return reflect.String, nil
	case "process.ancestors.envs_truncated":
		return reflect.Bool, nil
	case "process.ancestors.euid":
		return reflect.Int, nil
	case "process.ancestors.euser":
		return reflect.String, nil
	case "process.ancestors.file.change_time":
		return reflect.Int, nil
	case "process.ancestors.file.filesystem":
		return reflect.String, nil
	case "process.ancestors.file.gid":
		return reflect.Int, nil
	case "process.ancestors.file.group":
		return reflect.String, nil
	case "process.ancestors.file.in_upper_layer":
		return reflect.Bool, nil
	case "process.ancestors.file.inode":
		return reflect.Int, nil
	case "process.ancestors.file.mode":
		return reflect.Int, nil
	case "process.ancestors.file.modification_time":
		return reflect.Int, nil
	case "process.ancestors.file.mount_id":
		return reflect.Int, nil
	case "process.ancestors.file.name":
		return reflect.String, nil
	case "process.ancestors.file.name.length":
		return reflect.Int, nil
	case "process.ancestors.file.path":
		return reflect.String, nil
	case "process.ancestors.file.path.length":
		return reflect.Int, nil
	case "process.ancestors.file.rights":
		return reflect.Int, nil
	case "process.ancestors.file.uid":
		return reflect.Int, nil
	case "process.ancestors.file.user":
		return reflect.String, nil
	case "process.ancestors.fsgid":
		return reflect.Int, nil
	case "process.ancestors.fsgroup":
		return reflect.String, nil
	case "process.ancestors.fsuid":
		return reflect.Int, nil
	case "process.ancestors.fsuser":
		return reflect.String, nil
	case "process.ancestors.gid":
		return reflect.Int, nil
	case "process.ancestors.group":
		return reflect.String, nil
	case "process.ancestors.interpreter.file.change_time":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.filesystem":
		return reflect.String, nil
	case "process.ancestors.interpreter.file.gid":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.group":
		return reflect.String, nil
	case "process.ancestors.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "process.ancestors.interpreter.file.inode":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.mode":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.modification_time":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.mount_id":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.name":
		return reflect.String, nil
	case "process.ancestors.interpreter.file.name.length":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.path":
		return reflect.String, nil
	case "process.ancestors.interpreter.file.path.length":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.rights":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.uid":
		return reflect.Int, nil
	case "process.ancestors.interpreter.file.user":
		return reflect.String, nil
	case "process.ancestors.is_kworker":
		return reflect.Bool, nil
	case "process.ancestors.is_thread":
		return reflect.Bool, nil
	case "process.ancestors.pid":
		return reflect.Int, nil
	case "process.ancestors.ppid":
		return reflect.Int, nil
	case "process.ancestors.tid":
		return reflect.Int, nil
	case "process.ancestors.tty_name":
		return reflect.String, nil
	case "process.ancestors.uid":
		return reflect.Int, nil
	case "process.ancestors.user":
		return reflect.String, nil
	case "process.args":
		return reflect.String, nil
	case "process.args_flags":
		return reflect.String, nil
	case "process.args_options":
		return reflect.String, nil
	case "process.args_truncated":
		return reflect.Bool, nil
	case "process.argv":
		return reflect.String, nil
	case "process.argv0":
		return reflect.String, nil
	case "process.cap_effective":
		return reflect.Int, nil
	case "process.cap_permitted":
		return reflect.Int, nil
	case "process.comm":
		return reflect.String, nil
	case "process.container.id":
		return reflect.String, nil
	case "process.cookie":
		return reflect.Int, nil
	case "process.created_at":
		return reflect.Int, nil
	case "process.egid":
		return reflect.Int, nil
	case "process.egroup":
		return reflect.String, nil
	case "process.envp":
		return reflect.String, nil
	case "process.envs":
		return reflect.String, nil
	case "process.envs_truncated":
		return reflect.Bool, nil
	case "process.euid":
		return reflect.Int, nil
	case "process.euser":
		return reflect.String, nil
	case "process.file.change_time":
		return reflect.Int, nil
	case "process.file.filesystem":
		return reflect.String, nil
	case "process.file.gid":
		return reflect.Int, nil
	case "process.file.group":
		return reflect.String, nil
	case "process.file.in_upper_layer":
		return reflect.Bool, nil
	case "process.file.inode":
		return reflect.Int, nil
	case "process.file.mode":
		return reflect.Int, nil
	case "process.file.modification_time":
		return reflect.Int, nil
	case "process.file.mount_id":
		return reflect.Int, nil
	case "process.file.name":
		return reflect.String, nil
	case "process.file.name.length":
		return reflect.Int, nil
	case "process.file.path":
		return reflect.String, nil
	case "process.file.path.length":
		return reflect.Int, nil
	case "process.file.rights":
		return reflect.Int, nil
	case "process.file.uid":
		return reflect.Int, nil
	case "process.file.user":
		return reflect.String, nil
	case "process.fsgid":
		return reflect.Int, nil
	case "process.fsgroup":
		return reflect.String, nil
	case "process.fsuid":
		return reflect.Int, nil
	case "process.fsuser":
		return reflect.String, nil
	case "process.gid":
		return reflect.Int, nil
	case "process.group":
		return reflect.String, nil
	case "process.interpreter.file.change_time":
		return reflect.Int, nil
	case "process.interpreter.file.filesystem":
		return reflect.String, nil
	case "process.interpreter.file.gid":
		return reflect.Int, nil
	case "process.interpreter.file.group":
		return reflect.String, nil
	case "process.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "process.interpreter.file.inode":
		return reflect.Int, nil
	case "process.interpreter.file.mode":
		return reflect.Int, nil
	case "process.interpreter.file.modification_time":
		return reflect.Int, nil
	case "process.interpreter.file.mount_id":
		return reflect.Int, nil
	case "process.interpreter.file.name":
		return reflect.String, nil
	case "process.interpreter.file.name.length":
		return reflect.Int, nil
	case "process.interpreter.file.path":
		return reflect.String, nil
	case "process.interpreter.file.path.length":
		return reflect.Int, nil
	case "process.interpreter.file.rights":
		return reflect.Int, nil
	case "process.interpreter.file.uid":
		return reflect.Int, nil
	case "process.interpreter.file.user":
		return reflect.String, nil
	case "process.is_kworker":
		return reflect.Bool, nil
	case "process.is_thread":
		return reflect.Bool, nil
	case "process.parent.args":
		return reflect.String, nil
	case "process.parent.args_flags":
		return reflect.String, nil
	case "process.parent.args_options":
		return reflect.String, nil
	case "process.parent.args_truncated":
		return reflect.Bool, nil
	case "process.parent.argv":
		return reflect.String, nil
	case "process.parent.argv0":
		return reflect.String, nil
	case "process.parent.cap_effective":
		return reflect.Int, nil
	case "process.parent.cap_permitted":
		return reflect.Int, nil
	case "process.parent.comm":
		return reflect.String, nil
	case "process.parent.container.id":
		return reflect.String, nil
	case "process.parent.cookie":
		return reflect.Int, nil
	case "process.parent.created_at":
		return reflect.Int, nil
	case "process.parent.egid":
		return reflect.Int, nil
	case "process.parent.egroup":
		return reflect.String, nil
	case "process.parent.envp":
		return reflect.String, nil
	case "process.parent.envs":
		return reflect.String, nil
	case "process.parent.envs_truncated":
		return reflect.Bool, nil
	case "process.parent.euid":
		return reflect.Int, nil
	case "process.parent.euser":
		return reflect.String, nil
	case "process.parent.file.change_time":
		return reflect.Int, nil
	case "process.parent.file.filesystem":
		return reflect.String, nil
	case "process.parent.file.gid":
		return reflect.Int, nil
	case "process.parent.file.group":
		return reflect.String, nil
	case "process.parent.file.in_upper_layer":
		return reflect.Bool, nil
	case "process.parent.file.inode":
		return reflect.Int, nil
	case "process.parent.file.mode":
		return reflect.Int, nil
	case "process.parent.file.modification_time":
		return reflect.Int, nil
	case "process.parent.file.mount_id":
		return reflect.Int, nil
	case "process.parent.file.name":
		return reflect.String, nil
	case "process.parent.file.name.length":
		return reflect.Int, nil
	case "process.parent.file.path":
		return reflect.String, nil
	case "process.parent.file.path.length":
		return reflect.Int, nil
	case "process.parent.file.rights":
		return reflect.Int, nil
	case "process.parent.file.uid":
		return reflect.Int, nil
	case "process.parent.file.user":
		return reflect.String, nil
	case "process.parent.fsgid":
		return reflect.Int, nil
	case "process.parent.fsgroup":
		return reflect.String, nil
	case "process.parent.fsuid":
		return reflect.Int, nil
	case "process.parent.fsuser":
		return reflect.String, nil
	case "process.parent.gid":
		return reflect.Int, nil
	case "process.parent.group":
		return reflect.String, nil
	case "process.parent.interpreter.file.change_time":
		return reflect.Int, nil
	case "process.parent.interpreter.file.filesystem":
		return reflect.String, nil
	case "process.parent.interpreter.file.gid":
		return reflect.Int, nil
	case "process.parent.interpreter.file.group":
		return reflect.String, nil
	case "process.parent.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "process.parent.interpreter.file.inode":
		return reflect.Int, nil
	case "process.parent.interpreter.file.mode":
		return reflect.Int, nil
	case "process.parent.interpreter.file.modification_time":
		return reflect.Int, nil
	case "process.parent.interpreter.file.mount_id":
		return reflect.Int, nil
	case "process.parent.interpreter.file.name":
		return reflect.String, nil
	case "process.parent.interpreter.file.name.length":
		return reflect.Int, nil
	case "process.parent.interpreter.file.path":
		return reflect.String, nil
	case "process.parent.interpreter.file.path.length":
		return reflect.Int, nil
	case "process.parent.interpreter.file.rights":
		return reflect.Int, nil
	case "process.parent.interpreter.file.uid":
		return reflect.Int, nil
	case "process.parent.interpreter.file.user":
		return reflect.String, nil
	case "process.parent.is_kworker":
		return reflect.Bool, nil
	case "process.parent.is_thread":
		return reflect.Bool, nil
	case "process.parent.pid":
		return reflect.Int, nil
	case "process.parent.ppid":
		return reflect.Int, nil
	case "process.parent.tid":
		return reflect.Int, nil
	case "process.parent.tty_name":
		return reflect.String, nil
	case "process.parent.uid":
		return reflect.Int, nil
	case "process.parent.user":
		return reflect.String, nil
	case "process.pid":
		return reflect.Int, nil
	case "process.ppid":
		return reflect.Int, nil
	case "process.tid":
		return reflect.Int, nil
	case "process.tty_name":
		return reflect.String, nil
	case "process.uid":
		return reflect.Int, nil
	case "process.user":
		return reflect.String, nil
	case "removexattr.file.change_time":
		return reflect.Int, nil
	case "removexattr.file.destination.name":
		return reflect.String, nil
	case "removexattr.file.destination.namespace":
		return reflect.String, nil
	case "removexattr.file.filesystem":
		return reflect.String, nil
	case "removexattr.file.gid":
		return reflect.Int, nil
	case "removexattr.file.group":
		return reflect.String, nil
	case "removexattr.file.in_upper_layer":
		return reflect.Bool, nil
	case "removexattr.file.inode":
		return reflect.Int, nil
	case "removexattr.file.mode":
		return reflect.Int, nil
	case "removexattr.file.modification_time":
		return reflect.Int, nil
	case "removexattr.file.mount_id":
		return reflect.Int, nil
	case "removexattr.file.name":
		return reflect.String, nil
	case "removexattr.file.name.length":
		return reflect.Int, nil
	case "removexattr.file.path":
		return reflect.String, nil
	case "removexattr.file.path.length":
		return reflect.Int, nil
	case "removexattr.file.rights":
		return reflect.Int, nil
	case "removexattr.file.uid":
		return reflect.Int, nil
	case "removexattr.file.user":
		return reflect.String, nil
	case "removexattr.retval":
		return reflect.Int, nil
	case "rename.file.change_time":
		return reflect.Int, nil
	case "rename.file.destination.change_time":
		return reflect.Int, nil
	case "rename.file.destination.filesystem":
		return reflect.String, nil
	case "rename.file.destination.gid":
		return reflect.Int, nil
	case "rename.file.destination.group":
		return reflect.String, nil
	case "rename.file.destination.in_upper_layer":
		return reflect.Bool, nil
	case "rename.file.destination.inode":
		return reflect.Int, nil
	case "rename.file.destination.mode":
		return reflect.Int, nil
	case "rename.file.destination.modification_time":
		return reflect.Int, nil
	case "rename.file.destination.mount_id":
		return reflect.Int, nil
	case "rename.file.destination.name":
		return reflect.String, nil
	case "rename.file.destination.name.length":
		return reflect.Int, nil
	case "rename.file.destination.path":
		return reflect.String, nil
	case "rename.file.destination.path.length":
		return reflect.Int, nil
	case "rename.file.destination.rights":
		return reflect.Int, nil
	case "rename.file.destination.uid":
		return reflect.Int, nil
	case "rename.file.destination.user":
		return reflect.String, nil
	case "rename.file.filesystem":
		return reflect.String, nil
	case "rename.file.gid":
		return reflect.Int, nil
	case "rename.file.group":
		return reflect.String, nil
	case "rename.file.in_upper_layer":
		return reflect.Bool, nil
	case "rename.file.inode":
		return reflect.Int, nil
	case "rename.file.mode":
		return reflect.Int, nil
	case "rename.file.modification_time":
		return reflect.Int, nil
	case "rename.file.mount_id":
		return reflect.Int, nil
	case "rename.file.name":
		return reflect.String, nil
	case "rename.file.name.length":
		return reflect.Int, nil
	case "rename.file.path":
		return reflect.String, nil
	case "rename.file.path.length":
		return reflect.Int, nil
	case "rename.file.rights":
		return reflect.Int, nil
	case "rename.file.uid":
		return reflect.Int, nil
	case "rename.file.user":
		return reflect.String, nil
	case "rename.retval":
		return reflect.Int, nil
	case "rmdir.file.change_time":
		return reflect.Int, nil
	case "rmdir.file.filesystem":
		return reflect.String, nil
	case "rmdir.file.gid":
		return reflect.Int, nil
	case "rmdir.file.group":
		return reflect.String, nil
	case "rmdir.file.in_upper_layer":
		return reflect.Bool, nil
	case "rmdir.file.inode":
		return reflect.Int, nil
	case "rmdir.file.mode":
		return reflect.Int, nil
	case "rmdir.file.modification_time":
		return reflect.Int, nil
	case "rmdir.file.mount_id":
		return reflect.Int, nil
	case "rmdir.file.name":
		return reflect.String, nil
	case "rmdir.file.name.length":
		return reflect.Int, nil
	case "rmdir.file.path":
		return reflect.String, nil
	case "rmdir.file.path.length":
		return reflect.Int, nil
	case "rmdir.file.rights":
		return reflect.Int, nil
	case "rmdir.file.uid":
		return reflect.Int, nil
	case "rmdir.file.user":
		return reflect.String, nil
	case "rmdir.retval":
		return reflect.Int, nil
	case "setgid.egid":
		return reflect.Int, nil
	case "setgid.egroup":
		return reflect.String, nil
	case "setgid.fsgid":
		return reflect.Int, nil
	case "setgid.fsgroup":
		return reflect.String, nil
	case "setgid.gid":
		return reflect.Int, nil
	case "setgid.group":
		return reflect.String, nil
	case "setuid.euid":
		return reflect.Int, nil
	case "setuid.euser":
		return reflect.String, nil
	case "setuid.fsuid":
		return reflect.Int, nil
	case "setuid.fsuser":
		return reflect.String, nil
	case "setuid.uid":
		return reflect.Int, nil
	case "setuid.user":
		return reflect.String, nil
	case "setxattr.file.change_time":
		return reflect.Int, nil
	case "setxattr.file.destination.name":
		return reflect.String, nil
	case "setxattr.file.destination.namespace":
		return reflect.String, nil
	case "setxattr.file.filesystem":
		return reflect.String, nil
	case "setxattr.file.gid":
		return reflect.Int, nil
	case "setxattr.file.group":
		return reflect.String, nil
	case "setxattr.file.in_upper_layer":
		return reflect.Bool, nil
	case "setxattr.file.inode":
		return reflect.Int, nil
	case "setxattr.file.mode":
		return reflect.Int, nil
	case "setxattr.file.modification_time":
		return reflect.Int, nil
	case "setxattr.file.mount_id":
		return reflect.Int, nil
	case "setxattr.file.name":
		return reflect.String, nil
	case "setxattr.file.name.length":
		return reflect.Int, nil
	case "setxattr.file.path":
		return reflect.String, nil
	case "setxattr.file.path.length":
		return reflect.Int, nil
	case "setxattr.file.rights":
		return reflect.Int, nil
	case "setxattr.file.uid":
		return reflect.Int, nil
	case "setxattr.file.user":
		return reflect.String, nil
	case "setxattr.retval":
		return reflect.Int, nil
	case "signal.pid":
		return reflect.Int, nil
	case "signal.retval":
		return reflect.Int, nil
	case "signal.target.ancestors.args":
		return reflect.String, nil
	case "signal.target.ancestors.args_flags":
		return reflect.String, nil
	case "signal.target.ancestors.args_options":
		return reflect.String, nil
	case "signal.target.ancestors.args_truncated":
		return reflect.Bool, nil
	case "signal.target.ancestors.argv":
		return reflect.String, nil
	case "signal.target.ancestors.argv0":
		return reflect.String, nil
	case "signal.target.ancestors.cap_effective":
		return reflect.Int, nil
	case "signal.target.ancestors.cap_permitted":
		return reflect.Int, nil
	case "signal.target.ancestors.comm":
		return reflect.String, nil
	case "signal.target.ancestors.container.id":
		return reflect.String, nil
	case "signal.target.ancestors.cookie":
		return reflect.Int, nil
	case "signal.target.ancestors.created_at":
		return reflect.Int, nil
	case "signal.target.ancestors.egid":
		return reflect.Int, nil
	case "signal.target.ancestors.egroup":
		return reflect.String, nil
	case "signal.target.ancestors.envp":
		return reflect.String, nil
	case "signal.target.ancestors.envs":
		return reflect.String, nil
	case "signal.target.ancestors.envs_truncated":
		return reflect.Bool, nil
	case "signal.target.ancestors.euid":
		return reflect.Int, nil
	case "signal.target.ancestors.euser":
		return reflect.String, nil
	case "signal.target.ancestors.file.change_time":
		return reflect.Int, nil
	case "signal.target.ancestors.file.filesystem":
		return reflect.String, nil
	case "signal.target.ancestors.file.gid":
		return reflect.Int, nil
	case "signal.target.ancestors.file.group":
		return reflect.String, nil
	case "signal.target.ancestors.file.in_upper_layer":
		return reflect.Bool, nil
	case "signal.target.ancestors.file.inode":
		return reflect.Int, nil
	case "signal.target.ancestors.file.mode":
		return reflect.Int, nil
	case "signal.target.ancestors.file.modification_time":
		return reflect.Int, nil
	case "signal.target.ancestors.file.mount_id":
		return reflect.Int, nil
	case "signal.target.ancestors.file.name":
		return reflect.String, nil
	case "signal.target.ancestors.file.name.length":
		return reflect.Int, nil
	case "signal.target.ancestors.file.path":
		return reflect.String, nil
	case "signal.target.ancestors.file.path.length":
		return reflect.Int, nil
	case "signal.target.ancestors.file.rights":
		return reflect.Int, nil
	case "signal.target.ancestors.file.uid":
		return reflect.Int, nil
	case "signal.target.ancestors.file.user":
		return reflect.String, nil
	case "signal.target.ancestors.fsgid":
		return reflect.Int, nil
	case "signal.target.ancestors.fsgroup":
		return reflect.String, nil
	case "signal.target.ancestors.fsuid":
		return reflect.Int, nil
	case "signal.target.ancestors.fsuser":
		return reflect.String, nil
	case "signal.target.ancestors.gid":
		return reflect.Int, nil
	case "signal.target.ancestors.group":
		return reflect.String, nil
	case "signal.target.ancestors.interpreter.file.change_time":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.filesystem":
		return reflect.String, nil
	case "signal.target.ancestors.interpreter.file.gid":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.group":
		return reflect.String, nil
	case "signal.target.ancestors.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "signal.target.ancestors.interpreter.file.inode":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.mode":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.modification_time":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.mount_id":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.name":
		return reflect.String, nil
	case "signal.target.ancestors.interpreter.file.name.length":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.path":
		return reflect.String, nil
	case "signal.target.ancestors.interpreter.file.path.length":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.rights":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.uid":
		return reflect.Int, nil
	case "signal.target.ancestors.interpreter.file.user":
		return reflect.String, nil
	case "signal.target.ancestors.is_kworker":
		return reflect.Bool, nil
	case "signal.target.ancestors.is_thread":
		return reflect.Bool, nil
	case "signal.target.ancestors.pid":
		return reflect.Int, nil
	case "signal.target.ancestors.ppid":
		return reflect.Int, nil
	case "signal.target.ancestors.tid":
		return reflect.Int, nil
	case "signal.target.ancestors.tty_name":
		return reflect.String, nil
	case "signal.target.ancestors.uid":
		return reflect.Int, nil
	case "signal.target.ancestors.user":
		return reflect.String, nil
	case "signal.target.args":
		return reflect.String, nil
	case "signal.target.args_flags":
		return reflect.String, nil
	case "signal.target.args_options":
		return reflect.String, nil
	case "signal.target.args_truncated":
		return reflect.Bool, nil
	case "signal.target.argv":
		return reflect.String, nil
	case "signal.target.argv0":
		return reflect.String, nil
	case "signal.target.cap_effective":
		return reflect.Int, nil
	case "signal.target.cap_permitted":
		return reflect.Int, nil
	case "signal.target.comm":
		return reflect.String, nil
	case "signal.target.container.id":
		return reflect.String, nil
	case "signal.target.cookie":
		return reflect.Int, nil
	case "signal.target.created_at":
		return reflect.Int, nil
	case "signal.target.egid":
		return reflect.Int, nil
	case "signal.target.egroup":
		return reflect.String, nil
	case "signal.target.envp":
		return reflect.String, nil
	case "signal.target.envs":
		return reflect.String, nil
	case "signal.target.envs_truncated":
		return reflect.Bool, nil
	case "signal.target.euid":
		return reflect.Int, nil
	case "signal.target.euser":
		return reflect.String, nil
	case "signal.target.file.change_time":
		return reflect.Int, nil
	case "signal.target.file.filesystem":
		return reflect.String, nil
	case "signal.target.file.gid":
		return reflect.Int, nil
	case "signal.target.file.group":
		return reflect.String, nil
	case "signal.target.file.in_upper_layer":
		return reflect.Bool, nil
	case "signal.target.file.inode":
		return reflect.Int, nil
	case "signal.target.file.mode":
		return reflect.Int, nil
	case "signal.target.file.modification_time":
		return reflect.Int, nil
	case "signal.target.file.mount_id":
		return reflect.Int, nil
	case "signal.target.file.name":
		return reflect.String, nil
	case "signal.target.file.name.length":
		return reflect.Int, nil
	case "signal.target.file.path":
		return reflect.String, nil
	case "signal.target.file.path.length":
		return reflect.Int, nil
	case "signal.target.file.rights":
		return reflect.Int, nil
	case "signal.target.file.uid":
		return reflect.Int, nil
	case "signal.target.file.user":
		return reflect.String, nil
	case "signal.target.fsgid":
		return reflect.Int, nil
	case "signal.target.fsgroup":
		return reflect.String, nil
	case "signal.target.fsuid":
		return reflect.Int, nil
	case "signal.target.fsuser":
		return reflect.String, nil
	case "signal.target.gid":
		return reflect.Int, nil
	case "signal.target.group":
		return reflect.String, nil
	case "signal.target.interpreter.file.change_time":
		return reflect.Int, nil
	case "signal.target.interpreter.file.filesystem":
		return reflect.String, nil
	case "signal.target.interpreter.file.gid":
		return reflect.Int, nil
	case "signal.target.interpreter.file.group":
		return reflect.String, nil
	case "signal.target.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "signal.target.interpreter.file.inode":
		return reflect.Int, nil
	case "signal.target.interpreter.file.mode":
		return reflect.Int, nil
	case "signal.target.interpreter.file.modification_time":
		return reflect.Int, nil
	case "signal.target.interpreter.file.mount_id":
		return reflect.Int, nil
	case "signal.target.interpreter.file.name":
		return reflect.String, nil
	case "signal.target.interpreter.file.name.length":
		return reflect.Int, nil
	case "signal.target.interpreter.file.path":
		return reflect.String, nil
	case "signal.target.interpreter.file.path.length":
		return reflect.Int, nil
	case "signal.target.interpreter.file.rights":
		return reflect.Int, nil
	case "signal.target.interpreter.file.uid":
		return reflect.Int, nil
	case "signal.target.interpreter.file.user":
		return reflect.String, nil
	case "signal.target.is_kworker":
		return reflect.Bool, nil
	case "signal.target.is_thread":
		return reflect.Bool, nil
	case "signal.target.parent.args":
		return reflect.String, nil
	case "signal.target.parent.args_flags":
		return reflect.String, nil
	case "signal.target.parent.args_options":
		return reflect.String, nil
	case "signal.target.parent.args_truncated":
		return reflect.Bool, nil
	case "signal.target.parent.argv":
		return reflect.String, nil
	case "signal.target.parent.argv0":
		return reflect.String, nil
	case "signal.target.parent.cap_effective":
		return reflect.Int, nil
	case "signal.target.parent.cap_permitted":
		return reflect.Int, nil
	case "signal.target.parent.comm":
		return reflect.String, nil
	case "signal.target.parent.container.id":
		return reflect.String, nil
	case "signal.target.parent.cookie":
		return reflect.Int, nil
	case "signal.target.parent.created_at":
		return reflect.Int, nil
	case "signal.target.parent.egid":
		return reflect.Int, nil
	case "signal.target.parent.egroup":
		return reflect.String, nil
	case "signal.target.parent.envp":
		return reflect.String, nil
	case "signal.target.parent.envs":
		return reflect.String, nil
	case "signal.target.parent.envs_truncated":
		return reflect.Bool, nil
	case "signal.target.parent.euid":
		return reflect.Int, nil
	case "signal.target.parent.euser":
		return reflect.String, nil
	case "signal.target.parent.file.change_time":
		return reflect.Int, nil
	case "signal.target.parent.file.filesystem":
		return reflect.String, nil
	case "signal.target.parent.file.gid":
		return reflect.Int, nil
	case "signal.target.parent.file.group":
		return reflect.String, nil
	case "signal.target.parent.file.in_upper_layer":
		return reflect.Bool, nil
	case "signal.target.parent.file.inode":
		return reflect.Int, nil
	case "signal.target.parent.file.mode":
		return reflect.Int, nil
	case "signal.target.parent.file.modification_time":
		return reflect.Int, nil
	case "signal.target.parent.file.mount_id":
		return reflect.Int, nil
	case "signal.target.parent.file.name":
		return reflect.String, nil
	case "signal.target.parent.file.name.length":
		return reflect.Int, nil
	case "signal.target.parent.file.path":
		return reflect.String, nil
	case "signal.target.parent.file.path.length":
		return reflect.Int, nil
	case "signal.target.parent.file.rights":
		return reflect.Int, nil
	case "signal.target.parent.file.uid":
		return reflect.Int, nil
	case "signal.target.parent.file.user":
		return reflect.String, nil
	case "signal.target.parent.fsgid":
		return reflect.Int, nil
	case "signal.target.parent.fsgroup":
		return reflect.String, nil
	case "signal.target.parent.fsuid":
		return reflect.Int, nil
	case "signal.target.parent.fsuser":
		return reflect.String, nil
	case "signal.target.parent.gid":
		return reflect.Int, nil
	case "signal.target.parent.group":
		return reflect.String, nil
	case "signal.target.parent.interpreter.file.change_time":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.filesystem":
		return reflect.String, nil
	case "signal.target.parent.interpreter.file.gid":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.group":
		return reflect.String, nil
	case "signal.target.parent.interpreter.file.in_upper_layer":
		return reflect.Bool, nil
	case "signal.target.parent.interpreter.file.inode":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.mode":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.modification_time":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.mount_id":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.name":
		return reflect.String, nil
	case "signal.target.parent.interpreter.file.name.length":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.path":
		return reflect.String, nil
	case "signal.target.parent.interpreter.file.path.length":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.rights":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.uid":
		return reflect.Int, nil
	case "signal.target.parent.interpreter.file.user":
		return reflect.String, nil
	case "signal.target.parent.is_kworker":
		return reflect.Bool, nil
	case "signal.target.parent.is_thread":
		return reflect.Bool, nil
	case "signal.target.parent.pid":
		return reflect.Int, nil
	case "signal.target.parent.ppid":
		return reflect.Int, nil
	case "signal.target.parent.tid":
		return reflect.Int, nil
	case "signal.target.parent.tty_name":
		return reflect.String, nil
	case "signal.target.parent.uid":
		return reflect.Int, nil
	case "signal.target.parent.user":
		return reflect.String, nil
	case "signal.target.pid":
		return reflect.Int, nil
	case "signal.target.ppid":
		return reflect.Int, nil
	case "signal.target.tid":
		return reflect.Int, nil
	case "signal.target.tty_name":
		return reflect.String, nil
	case "signal.target.uid":
		return reflect.Int, nil
	case "signal.target.user":
		return reflect.String, nil
	case "signal.type":
		return reflect.Int, nil
	case "splice.file.change_time":
		return reflect.Int, nil
	case "splice.file.filesystem":
		return reflect.String, nil
	case "splice.file.gid":
		return reflect.Int, nil
	case "splice.file.group":
		return reflect.String, nil
	case "splice.file.in_upper_layer":
		return reflect.Bool, nil
	case "splice.file.inode":
		return reflect.Int, nil
	case "splice.file.mode":
		return reflect.Int, nil
	case "splice.file.modification_time":
		return reflect.Int, nil
	case "splice.file.mount_id":
		return reflect.Int, nil
	case "splice.file.name":
		return reflect.String, nil
	case "splice.file.name.length":
		return reflect.Int, nil
	case "splice.file.path":
		return reflect.String, nil
	case "splice.file.path.length":
		return reflect.Int, nil
	case "splice.file.rights":
		return reflect.Int, nil
	case "splice.file.uid":
		return reflect.Int, nil
	case "splice.file.user":
		return reflect.String, nil
	case "splice.pipe_entry_flag":
		return reflect.Int, nil
	case "splice.pipe_exit_flag":
		return reflect.Int, nil
	case "splice.retval":
		return reflect.Int, nil
	case "unlink.file.change_time":
		return reflect.Int, nil
	case "unlink.file.filesystem":
		return reflect.String, nil
	case "unlink.file.gid":
		return reflect.Int, nil
	case "unlink.file.group":
		return reflect.String, nil
	case "unlink.file.in_upper_layer":
		return reflect.Bool, nil
	case "unlink.file.inode":
		return reflect.Int, nil
	case "unlink.file.mode":
		return reflect.Int, nil
	case "unlink.file.modification_time":
		return reflect.Int, nil
	case "unlink.file.mount_id":
		return reflect.Int, nil
	case "unlink.file.name":
		return reflect.String, nil
	case "unlink.file.name.length":
		return reflect.Int, nil
	case "unlink.file.path":
		return reflect.String, nil
	case "unlink.file.path.length":
		return reflect.Int, nil
	case "unlink.file.rights":
		return reflect.Int, nil
	case "unlink.file.uid":
		return reflect.Int, nil
	case "unlink.file.user":
		return reflect.String, nil
	case "unlink.flags":
		return reflect.Int, nil
	case "unlink.retval":
		return reflect.Int, nil
	case "utimes.file.change_time":
		return reflect.Int, nil
	case "utimes.file.filesystem":
		return reflect.String, nil
	case "utimes.file.gid":
		return reflect.Int, nil
	case "utimes.file.group":
		return reflect.String, nil
	case "utimes.file.in_upper_layer":
		return reflect.Bool, nil
	case "utimes.file.inode":
		return reflect.Int, nil
	case "utimes.file.mode":
		return reflect.Int, nil
	case "utimes.file.modification_time":
		return reflect.Int, nil
	case "utimes.file.mount_id":
		return reflect.Int, nil
	case "utimes.file.name":
		return reflect.String, nil
	case "utimes.file.name.length":
		return reflect.Int, nil
	case "utimes.file.path":
		return reflect.String, nil
	case "utimes.file.path.length":
		return reflect.Int, nil
	case "utimes.file.rights":
		return reflect.Int, nil
	case "utimes.file.uid":
		return reflect.Int, nil
	case "utimes.file.user":
		return reflect.String, nil
	case "utimes.retval":
		return reflect.Int, nil
	}
	return reflect.Invalid, &eval.ErrFieldNotFound{Field: field}
}
func (ev *Event) SetFieldValue(field eval.Field, value interface{}) error {
	switch field {
	case "async":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Async"}
		}
		ev.Async = rv
		return nil
	case "bind.addr.family":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Bind.AddrFamily"}
		}
		ev.Bind.AddrFamily = uint16(rv)
		return nil
	case "bind.addr.ip":
		rv, ok := value.(net.IPNet)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Bind.Addr.IPNet"}
		}
		ev.Bind.Addr.IPNet = rv
		return nil
	case "bind.addr.port":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Bind.Addr.Port"}
		}
		ev.Bind.Addr.Port = uint16(rv)
		return nil
	case "bind.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Bind.SyscallEvent.Retval"}
		}
		ev.Bind.SyscallEvent.Retval = int64(rv)
		return nil
	case "capset.cap_effective":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Capset.CapEffective"}
		}
		ev.Capset.CapEffective = uint64(rv)
		return nil
	case "capset.cap_permitted":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Capset.CapPermitted"}
		}
		ev.Capset.CapPermitted = uint64(rv)
		return nil
	case "chmod.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.CTime"}
		}
		ev.Chmod.File.FileFields.CTime = uint64(rv)
		return nil
	case "chmod.file.destination.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.Mode"}
		}
		ev.Chmod.Mode = uint32(rv)
		return nil
	case "chmod.file.destination.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.Mode"}
		}
		ev.Chmod.Mode = uint32(rv)
		return nil
	case "chmod.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.Filesystem"}
		}
		ev.Chmod.File.Filesystem = rv
		return nil
	case "chmod.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.GID"}
		}
		ev.Chmod.File.FileFields.GID = uint32(rv)
		return nil
	case "chmod.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.Group"}
		}
		ev.Chmod.File.FileFields.Group = rv
		return nil
	case "chmod.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.InUpperLayer"}
		}
		ev.Chmod.File.FileFields.InUpperLayer = rv
		return nil
	case "chmod.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.Inode"}
		}
		ev.Chmod.File.FileFields.Inode = uint64(rv)
		return nil
	case "chmod.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.Mode"}
		}
		ev.Chmod.File.FileFields.Mode = uint16(rv)
		return nil
	case "chmod.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.MTime"}
		}
		ev.Chmod.File.FileFields.MTime = uint64(rv)
		return nil
	case "chmod.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.MountID"}
		}
		ev.Chmod.File.FileFields.MountID = uint32(rv)
		return nil
	case "chmod.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.BasenameStr"}
		}
		ev.Chmod.File.BasenameStr = rv
		return nil
	case "chmod.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "chmod.file.name.length"}
	case "chmod.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.PathnameStr"}
		}
		ev.Chmod.File.PathnameStr = rv
		return nil
	case "chmod.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "chmod.file.path.length"}
	case "chmod.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.Mode"}
		}
		ev.Chmod.File.FileFields.Mode = uint16(rv)
		return nil
	case "chmod.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.UID"}
		}
		ev.Chmod.File.FileFields.UID = uint32(rv)
		return nil
	case "chmod.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.File.FileFields.User"}
		}
		ev.Chmod.File.FileFields.User = rv
		return nil
	case "chmod.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chmod.SyscallEvent.Retval"}
		}
		ev.Chmod.SyscallEvent.Retval = int64(rv)
		return nil
	case "chown.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.CTime"}
		}
		ev.Chown.File.FileFields.CTime = uint64(rv)
		return nil
	case "chown.file.destination.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.GID"}
		}
		ev.Chown.GID = int64(rv)
		return nil
	case "chown.file.destination.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.Group"}
		}
		ev.Chown.Group = rv
		return nil
	case "chown.file.destination.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.UID"}
		}
		ev.Chown.UID = int64(rv)
		return nil
	case "chown.file.destination.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.User"}
		}
		ev.Chown.User = rv
		return nil
	case "chown.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.Filesystem"}
		}
		ev.Chown.File.Filesystem = rv
		return nil
	case "chown.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.GID"}
		}
		ev.Chown.File.FileFields.GID = uint32(rv)
		return nil
	case "chown.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.Group"}
		}
		ev.Chown.File.FileFields.Group = rv
		return nil
	case "chown.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.InUpperLayer"}
		}
		ev.Chown.File.FileFields.InUpperLayer = rv
		return nil
	case "chown.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.Inode"}
		}
		ev.Chown.File.FileFields.Inode = uint64(rv)
		return nil
	case "chown.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.Mode"}
		}
		ev.Chown.File.FileFields.Mode = uint16(rv)
		return nil
	case "chown.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.MTime"}
		}
		ev.Chown.File.FileFields.MTime = uint64(rv)
		return nil
	case "chown.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.MountID"}
		}
		ev.Chown.File.FileFields.MountID = uint32(rv)
		return nil
	case "chown.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.BasenameStr"}
		}
		ev.Chown.File.BasenameStr = rv
		return nil
	case "chown.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "chown.file.name.length"}
	case "chown.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.PathnameStr"}
		}
		ev.Chown.File.PathnameStr = rv
		return nil
	case "chown.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "chown.file.path.length"}
	case "chown.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.Mode"}
		}
		ev.Chown.File.FileFields.Mode = uint16(rv)
		return nil
	case "chown.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.UID"}
		}
		ev.Chown.File.FileFields.UID = uint32(rv)
		return nil
	case "chown.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.File.FileFields.User"}
		}
		ev.Chown.File.FileFields.User = rv
		return nil
	case "chown.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Chown.SyscallEvent.Retval"}
		}
		ev.Chown.SyscallEvent.Retval = int64(rv)
		return nil
	case "container.id":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ContainerContext.ID"}
		}
		ev.ContainerContext.ID = rv
		return nil
	case "container.tags":
		switch rv := value.(type) {
		case string:
			ev.ContainerContext.Tags = append(ev.ContainerContext.Tags, rv)
		case []string:
			ev.ContainerContext.Tags = append(ev.ContainerContext.Tags, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ContainerContext.Tags"}
		}
		return nil
	case "dns.id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "DNS.ID"}
		}
		ev.DNS.ID = uint16(rv)
		return nil
	case "dns.question.class":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "DNS.Class"}
		}
		ev.DNS.Class = uint16(rv)
		return nil
	case "dns.question.count":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "DNS.Count"}
		}
		ev.DNS.Count = uint16(rv)
		return nil
	case "dns.question.length":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "DNS.Size"}
		}
		ev.DNS.Size = uint16(rv)
		return nil
	case "dns.question.name":
		return &eval.ErrFieldReadOnly{Field: "dns.question.name"}
	case "dns.question.type":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "DNS.Type"}
		}
		ev.DNS.Type = uint16(rv)
		return nil
	case "exec.args":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Args"}
		}
		ev.Exec.Process.Args = rv
		return nil
	case "exec.args_flags":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exec.Process.Argv = append(ev.Exec.Process.Argv, rv)
		case []string:
			ev.Exec.Process.Argv = append(ev.Exec.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Argv"}
		}
		return nil
	case "exec.args_options":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exec.Process.Argv = append(ev.Exec.Process.Argv, rv)
		case []string:
			ev.Exec.Process.Argv = append(ev.Exec.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Argv"}
		}
		return nil
	case "exec.args_truncated":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.ArgsTruncated"}
		}
		ev.Exec.Process.ArgsTruncated = rv
		return nil
	case "exec.argv":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exec.Process.Argv = append(ev.Exec.Process.Argv, rv)
		case []string:
			ev.Exec.Process.Argv = append(ev.Exec.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Argv"}
		}
		return nil
	case "exec.argv0":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Argv0"}
		}
		ev.Exec.Process.Argv0 = rv
		return nil
	case "exec.cap_effective":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.CapEffective"}
		}
		ev.Exec.Process.Credentials.CapEffective = uint64(rv)
		return nil
	case "exec.cap_permitted":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.CapPermitted"}
		}
		ev.Exec.Process.Credentials.CapPermitted = uint64(rv)
		return nil
	case "exec.comm":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Comm"}
		}
		ev.Exec.Process.Comm = rv
		return nil
	case "exec.container.id":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.ContainerID"}
		}
		ev.Exec.Process.ContainerID = rv
		return nil
	case "exec.cookie":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Cookie"}
		}
		ev.Exec.Process.Cookie = uint32(rv)
		return nil
	case "exec.created_at":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.CreatedAt"}
		}
		ev.Exec.Process.CreatedAt = uint64(rv)
		return nil
	case "exec.egid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.EGID"}
		}
		ev.Exec.Process.Credentials.EGID = uint32(rv)
		return nil
	case "exec.egroup":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.EGroup"}
		}
		ev.Exec.Process.Credentials.EGroup = rv
		return nil
	case "exec.envp":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exec.Process.Envp = append(ev.Exec.Process.Envp, rv)
		case []string:
			ev.Exec.Process.Envp = append(ev.Exec.Process.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Envp"}
		}
		return nil
	case "exec.envs":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exec.Process.Envs = append(ev.Exec.Process.Envs, rv)
		case []string:
			ev.Exec.Process.Envs = append(ev.Exec.Process.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Envs"}
		}
		return nil
	case "exec.envs_truncated":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.EnvsTruncated"}
		}
		ev.Exec.Process.EnvsTruncated = rv
		return nil
	case "exec.euid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.EUID"}
		}
		ev.Exec.Process.Credentials.EUID = uint32(rv)
		return nil
	case "exec.euser":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.EUser"}
		}
		ev.Exec.Process.Credentials.EUser = rv
		return nil
	case "exec.file.change_time":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.CTime"}
		}
		ev.Exec.Process.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "exec.file.filesystem":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.Filesystem"}
		}
		ev.Exec.Process.FileEvent.Filesystem = rv
		return nil
	case "exec.file.gid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.GID"}
		}
		ev.Exec.Process.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "exec.file.group":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.Group"}
		}
		ev.Exec.Process.FileEvent.FileFields.Group = rv
		return nil
	case "exec.file.in_upper_layer":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Exec.Process.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "exec.file.inode":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.Inode"}
		}
		ev.Exec.Process.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "exec.file.mode":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.Mode"}
		}
		ev.Exec.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exec.file.modification_time":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.MTime"}
		}
		ev.Exec.Process.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "exec.file.mount_id":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.MountID"}
		}
		ev.Exec.Process.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "exec.file.name":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.BasenameStr"}
		}
		ev.Exec.Process.FileEvent.BasenameStr = rv
		return nil
	case "exec.file.name.length":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exec.file.name.length"}
	case "exec.file.path":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.PathnameStr"}
		}
		ev.Exec.Process.FileEvent.PathnameStr = rv
		return nil
	case "exec.file.path.length":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exec.file.path.length"}
	case "exec.file.rights":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.Mode"}
		}
		ev.Exec.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exec.file.uid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.UID"}
		}
		ev.Exec.Process.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "exec.file.user":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.FileEvent.FileFields.User"}
		}
		ev.Exec.Process.FileEvent.FileFields.User = rv
		return nil
	case "exec.fsgid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.FSGID"}
		}
		ev.Exec.Process.Credentials.FSGID = uint32(rv)
		return nil
	case "exec.fsgroup":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.FSGroup"}
		}
		ev.Exec.Process.Credentials.FSGroup = rv
		return nil
	case "exec.fsuid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.FSUID"}
		}
		ev.Exec.Process.Credentials.FSUID = uint32(rv)
		return nil
	case "exec.fsuser":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.FSUser"}
		}
		ev.Exec.Process.Credentials.FSUser = rv
		return nil
	case "exec.gid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.GID"}
		}
		ev.Exec.Process.Credentials.GID = uint32(rv)
		return nil
	case "exec.group":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.Group"}
		}
		ev.Exec.Process.Credentials.Group = rv
		return nil
	case "exec.interpreter.file.change_time":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "exec.interpreter.file.filesystem":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "exec.interpreter.file.gid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "exec.interpreter.file.group":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "exec.interpreter.file.in_upper_layer":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "exec.interpreter.file.inode":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "exec.interpreter.file.mode":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exec.interpreter.file.modification_time":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "exec.interpreter.file.mount_id":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "exec.interpreter.file.name":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "exec.interpreter.file.name.length":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exec.interpreter.file.name.length"}
	case "exec.interpreter.file.path":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "exec.interpreter.file.path.length":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exec.interpreter.file.path.length"}
	case "exec.interpreter.file.rights":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exec.interpreter.file.uid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "exec.interpreter.file.user":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.Exec.Process.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "exec.is_kworker":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.PIDContext.IsKworker"}
		}
		ev.Exec.Process.PIDContext.IsKworker = rv
		return nil
	case "exec.is_thread":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.IsThread"}
		}
		ev.Exec.Process.IsThread = rv
		return nil
	case "exec.pid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.PIDContext.Pid"}
		}
		ev.Exec.Process.PIDContext.Pid = uint32(rv)
		return nil
	case "exec.ppid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.PPid"}
		}
		ev.Exec.Process.PPid = uint32(rv)
		return nil
	case "exec.tid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.PIDContext.Tid"}
		}
		ev.Exec.Process.PIDContext.Tid = uint32(rv)
		return nil
	case "exec.tty_name":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.TTYName"}
		}
		ev.Exec.Process.TTYName = rv
		return nil
	case "exec.uid":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.UID"}
		}
		ev.Exec.Process.Credentials.UID = uint32(rv)
		return nil
	case "exec.user":
		if ev.Exec.Process == nil {
			ev.Exec.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exec.Process.Credentials.User"}
		}
		ev.Exec.Process.Credentials.User = rv
		return nil
	case "exit.args":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Args"}
		}
		ev.Exit.Process.Args = rv
		return nil
	case "exit.args_flags":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exit.Process.Argv = append(ev.Exit.Process.Argv, rv)
		case []string:
			ev.Exit.Process.Argv = append(ev.Exit.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Argv"}
		}
		return nil
	case "exit.args_options":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exit.Process.Argv = append(ev.Exit.Process.Argv, rv)
		case []string:
			ev.Exit.Process.Argv = append(ev.Exit.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Argv"}
		}
		return nil
	case "exit.args_truncated":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.ArgsTruncated"}
		}
		ev.Exit.Process.ArgsTruncated = rv
		return nil
	case "exit.argv":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exit.Process.Argv = append(ev.Exit.Process.Argv, rv)
		case []string:
			ev.Exit.Process.Argv = append(ev.Exit.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Argv"}
		}
		return nil
	case "exit.argv0":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Argv0"}
		}
		ev.Exit.Process.Argv0 = rv
		return nil
	case "exit.cap_effective":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.CapEffective"}
		}
		ev.Exit.Process.Credentials.CapEffective = uint64(rv)
		return nil
	case "exit.cap_permitted":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.CapPermitted"}
		}
		ev.Exit.Process.Credentials.CapPermitted = uint64(rv)
		return nil
	case "exit.cause":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Cause"}
		}
		ev.Exit.Cause = uint32(rv)
		return nil
	case "exit.code":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Code"}
		}
		ev.Exit.Code = uint32(rv)
		return nil
	case "exit.comm":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Comm"}
		}
		ev.Exit.Process.Comm = rv
		return nil
	case "exit.container.id":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.ContainerID"}
		}
		ev.Exit.Process.ContainerID = rv
		return nil
	case "exit.cookie":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Cookie"}
		}
		ev.Exit.Process.Cookie = uint32(rv)
		return nil
	case "exit.created_at":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.CreatedAt"}
		}
		ev.Exit.Process.CreatedAt = uint64(rv)
		return nil
	case "exit.egid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.EGID"}
		}
		ev.Exit.Process.Credentials.EGID = uint32(rv)
		return nil
	case "exit.egroup":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.EGroup"}
		}
		ev.Exit.Process.Credentials.EGroup = rv
		return nil
	case "exit.envp":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exit.Process.Envp = append(ev.Exit.Process.Envp, rv)
		case []string:
			ev.Exit.Process.Envp = append(ev.Exit.Process.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Envp"}
		}
		return nil
	case "exit.envs":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Exit.Process.Envs = append(ev.Exit.Process.Envs, rv)
		case []string:
			ev.Exit.Process.Envs = append(ev.Exit.Process.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Envs"}
		}
		return nil
	case "exit.envs_truncated":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.EnvsTruncated"}
		}
		ev.Exit.Process.EnvsTruncated = rv
		return nil
	case "exit.euid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.EUID"}
		}
		ev.Exit.Process.Credentials.EUID = uint32(rv)
		return nil
	case "exit.euser":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.EUser"}
		}
		ev.Exit.Process.Credentials.EUser = rv
		return nil
	case "exit.file.change_time":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.CTime"}
		}
		ev.Exit.Process.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "exit.file.filesystem":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.Filesystem"}
		}
		ev.Exit.Process.FileEvent.Filesystem = rv
		return nil
	case "exit.file.gid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.GID"}
		}
		ev.Exit.Process.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "exit.file.group":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.Group"}
		}
		ev.Exit.Process.FileEvent.FileFields.Group = rv
		return nil
	case "exit.file.in_upper_layer":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Exit.Process.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "exit.file.inode":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.Inode"}
		}
		ev.Exit.Process.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "exit.file.mode":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.Mode"}
		}
		ev.Exit.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exit.file.modification_time":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.MTime"}
		}
		ev.Exit.Process.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "exit.file.mount_id":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.MountID"}
		}
		ev.Exit.Process.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "exit.file.name":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.BasenameStr"}
		}
		ev.Exit.Process.FileEvent.BasenameStr = rv
		return nil
	case "exit.file.name.length":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exit.file.name.length"}
	case "exit.file.path":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.PathnameStr"}
		}
		ev.Exit.Process.FileEvent.PathnameStr = rv
		return nil
	case "exit.file.path.length":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exit.file.path.length"}
	case "exit.file.rights":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.Mode"}
		}
		ev.Exit.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exit.file.uid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.UID"}
		}
		ev.Exit.Process.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "exit.file.user":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.FileEvent.FileFields.User"}
		}
		ev.Exit.Process.FileEvent.FileFields.User = rv
		return nil
	case "exit.fsgid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.FSGID"}
		}
		ev.Exit.Process.Credentials.FSGID = uint32(rv)
		return nil
	case "exit.fsgroup":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.FSGroup"}
		}
		ev.Exit.Process.Credentials.FSGroup = rv
		return nil
	case "exit.fsuid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.FSUID"}
		}
		ev.Exit.Process.Credentials.FSUID = uint32(rv)
		return nil
	case "exit.fsuser":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.FSUser"}
		}
		ev.Exit.Process.Credentials.FSUser = rv
		return nil
	case "exit.gid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.GID"}
		}
		ev.Exit.Process.Credentials.GID = uint32(rv)
		return nil
	case "exit.group":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.Group"}
		}
		ev.Exit.Process.Credentials.Group = rv
		return nil
	case "exit.interpreter.file.change_time":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "exit.interpreter.file.filesystem":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "exit.interpreter.file.gid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "exit.interpreter.file.group":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "exit.interpreter.file.in_upper_layer":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "exit.interpreter.file.inode":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "exit.interpreter.file.mode":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exit.interpreter.file.modification_time":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "exit.interpreter.file.mount_id":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "exit.interpreter.file.name":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "exit.interpreter.file.name.length":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exit.interpreter.file.name.length"}
	case "exit.interpreter.file.path":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "exit.interpreter.file.path.length":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "exit.interpreter.file.path.length"}
	case "exit.interpreter.file.rights":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "exit.interpreter.file.uid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "exit.interpreter.file.user":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.Exit.Process.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "exit.is_kworker":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.PIDContext.IsKworker"}
		}
		ev.Exit.Process.PIDContext.IsKworker = rv
		return nil
	case "exit.is_thread":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.IsThread"}
		}
		ev.Exit.Process.IsThread = rv
		return nil
	case "exit.pid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.PIDContext.Pid"}
		}
		ev.Exit.Process.PIDContext.Pid = uint32(rv)
		return nil
	case "exit.ppid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.PPid"}
		}
		ev.Exit.Process.PPid = uint32(rv)
		return nil
	case "exit.tid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.PIDContext.Tid"}
		}
		ev.Exit.Process.PIDContext.Tid = uint32(rv)
		return nil
	case "exit.tty_name":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.TTYName"}
		}
		ev.Exit.Process.TTYName = rv
		return nil
	case "exit.uid":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.UID"}
		}
		ev.Exit.Process.Credentials.UID = uint32(rv)
		return nil
	case "exit.user":
		if ev.Exit.Process == nil {
			ev.Exit.Process = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Exit.Process.Credentials.User"}
		}
		ev.Exit.Process.Credentials.User = rv
		return nil
	case "link.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.CTime"}
		}
		ev.Link.Source.FileFields.CTime = uint64(rv)
		return nil
	case "link.file.destination.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.CTime"}
		}
		ev.Link.Target.FileFields.CTime = uint64(rv)
		return nil
	case "link.file.destination.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.Filesystem"}
		}
		ev.Link.Target.Filesystem = rv
		return nil
	case "link.file.destination.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.GID"}
		}
		ev.Link.Target.FileFields.GID = uint32(rv)
		return nil
	case "link.file.destination.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.Group"}
		}
		ev.Link.Target.FileFields.Group = rv
		return nil
	case "link.file.destination.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.InUpperLayer"}
		}
		ev.Link.Target.FileFields.InUpperLayer = rv
		return nil
	case "link.file.destination.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.Inode"}
		}
		ev.Link.Target.FileFields.Inode = uint64(rv)
		return nil
	case "link.file.destination.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.Mode"}
		}
		ev.Link.Target.FileFields.Mode = uint16(rv)
		return nil
	case "link.file.destination.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.MTime"}
		}
		ev.Link.Target.FileFields.MTime = uint64(rv)
		return nil
	case "link.file.destination.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.MountID"}
		}
		ev.Link.Target.FileFields.MountID = uint32(rv)
		return nil
	case "link.file.destination.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.BasenameStr"}
		}
		ev.Link.Target.BasenameStr = rv
		return nil
	case "link.file.destination.name.length":
		return &eval.ErrFieldReadOnly{Field: "link.file.destination.name.length"}
	case "link.file.destination.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.PathnameStr"}
		}
		ev.Link.Target.PathnameStr = rv
		return nil
	case "link.file.destination.path.length":
		return &eval.ErrFieldReadOnly{Field: "link.file.destination.path.length"}
	case "link.file.destination.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.Mode"}
		}
		ev.Link.Target.FileFields.Mode = uint16(rv)
		return nil
	case "link.file.destination.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.UID"}
		}
		ev.Link.Target.FileFields.UID = uint32(rv)
		return nil
	case "link.file.destination.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Target.FileFields.User"}
		}
		ev.Link.Target.FileFields.User = rv
		return nil
	case "link.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.Filesystem"}
		}
		ev.Link.Source.Filesystem = rv
		return nil
	case "link.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.GID"}
		}
		ev.Link.Source.FileFields.GID = uint32(rv)
		return nil
	case "link.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.Group"}
		}
		ev.Link.Source.FileFields.Group = rv
		return nil
	case "link.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.InUpperLayer"}
		}
		ev.Link.Source.FileFields.InUpperLayer = rv
		return nil
	case "link.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.Inode"}
		}
		ev.Link.Source.FileFields.Inode = uint64(rv)
		return nil
	case "link.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.Mode"}
		}
		ev.Link.Source.FileFields.Mode = uint16(rv)
		return nil
	case "link.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.MTime"}
		}
		ev.Link.Source.FileFields.MTime = uint64(rv)
		return nil
	case "link.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.MountID"}
		}
		ev.Link.Source.FileFields.MountID = uint32(rv)
		return nil
	case "link.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.BasenameStr"}
		}
		ev.Link.Source.BasenameStr = rv
		return nil
	case "link.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "link.file.name.length"}
	case "link.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.PathnameStr"}
		}
		ev.Link.Source.PathnameStr = rv
		return nil
	case "link.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "link.file.path.length"}
	case "link.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.Mode"}
		}
		ev.Link.Source.FileFields.Mode = uint16(rv)
		return nil
	case "link.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.UID"}
		}
		ev.Link.Source.FileFields.UID = uint32(rv)
		return nil
	case "link.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.Source.FileFields.User"}
		}
		ev.Link.Source.FileFields.User = rv
		return nil
	case "link.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Link.SyscallEvent.Retval"}
		}
		ev.Link.SyscallEvent.Retval = int64(rv)
		return nil
	case "mkdir.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.CTime"}
		}
		ev.Mkdir.File.FileFields.CTime = uint64(rv)
		return nil
	case "mkdir.file.destination.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.Mode"}
		}
		ev.Mkdir.Mode = uint32(rv)
		return nil
	case "mkdir.file.destination.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.Mode"}
		}
		ev.Mkdir.Mode = uint32(rv)
		return nil
	case "mkdir.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.Filesystem"}
		}
		ev.Mkdir.File.Filesystem = rv
		return nil
	case "mkdir.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.GID"}
		}
		ev.Mkdir.File.FileFields.GID = uint32(rv)
		return nil
	case "mkdir.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.Group"}
		}
		ev.Mkdir.File.FileFields.Group = rv
		return nil
	case "mkdir.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.InUpperLayer"}
		}
		ev.Mkdir.File.FileFields.InUpperLayer = rv
		return nil
	case "mkdir.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.Inode"}
		}
		ev.Mkdir.File.FileFields.Inode = uint64(rv)
		return nil
	case "mkdir.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.Mode"}
		}
		ev.Mkdir.File.FileFields.Mode = uint16(rv)
		return nil
	case "mkdir.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.MTime"}
		}
		ev.Mkdir.File.FileFields.MTime = uint64(rv)
		return nil
	case "mkdir.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.MountID"}
		}
		ev.Mkdir.File.FileFields.MountID = uint32(rv)
		return nil
	case "mkdir.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.BasenameStr"}
		}
		ev.Mkdir.File.BasenameStr = rv
		return nil
	case "mkdir.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "mkdir.file.name.length"}
	case "mkdir.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.PathnameStr"}
		}
		ev.Mkdir.File.PathnameStr = rv
		return nil
	case "mkdir.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "mkdir.file.path.length"}
	case "mkdir.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.Mode"}
		}
		ev.Mkdir.File.FileFields.Mode = uint16(rv)
		return nil
	case "mkdir.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.UID"}
		}
		ev.Mkdir.File.FileFields.UID = uint32(rv)
		return nil
	case "mkdir.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.File.FileFields.User"}
		}
		ev.Mkdir.File.FileFields.User = rv
		return nil
	case "mkdir.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mkdir.SyscallEvent.Retval"}
		}
		ev.Mkdir.SyscallEvent.Retval = int64(rv)
		return nil
	case "mount.fs_type":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mount.Mount.FSType"}
		}
		ev.Mount.Mount.FSType = rv
		return nil
	case "mount.mountpoint.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mount.MountPointPath"}
		}
		ev.Mount.MountPointPath = rv
		return nil
	case "mount.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mount.SyscallEvent.Retval"}
		}
		ev.Mount.SyscallEvent.Retval = int64(rv)
		return nil
	case "mount.source.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Mount.MountSourcePath"}
		}
		ev.Mount.MountSourcePath = rv
		return nil
	case "network.destination.ip":
		rv, ok := value.(net.IPNet)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.Destination.IPNet"}
		}
		ev.NetworkContext.Destination.IPNet = rv
		return nil
	case "network.destination.port":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.Destination.Port"}
		}
		ev.NetworkContext.Destination.Port = uint16(rv)
		return nil
	case "network.device.ifindex":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.Device.IfIndex"}
		}
		ev.NetworkContext.Device.IfIndex = uint32(rv)
		return nil
	case "network.device.ifname":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.Device.IfName"}
		}
		ev.NetworkContext.Device.IfName = rv
		return nil
	case "network.l3_protocol":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.L3Protocol"}
		}
		ev.NetworkContext.L3Protocol = uint16(rv)
		return nil
	case "network.l4_protocol":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.L4Protocol"}
		}
		ev.NetworkContext.L4Protocol = uint16(rv)
		return nil
	case "network.size":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.Size"}
		}
		ev.NetworkContext.Size = uint32(rv)
		return nil
	case "network.source.ip":
		rv, ok := value.(net.IPNet)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.Source.IPNet"}
		}
		ev.NetworkContext.Source.IPNet = rv
		return nil
	case "network.source.port":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "NetworkContext.Source.Port"}
		}
		ev.NetworkContext.Source.Port = uint16(rv)
		return nil
	case "open.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.CTime"}
		}
		ev.Open.File.FileFields.CTime = uint64(rv)
		return nil
	case "open.file.destination.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.Mode"}
		}
		ev.Open.Mode = uint32(rv)
		return nil
	case "open.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.Filesystem"}
		}
		ev.Open.File.Filesystem = rv
		return nil
	case "open.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.GID"}
		}
		ev.Open.File.FileFields.GID = uint32(rv)
		return nil
	case "open.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.Group"}
		}
		ev.Open.File.FileFields.Group = rv
		return nil
	case "open.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.InUpperLayer"}
		}
		ev.Open.File.FileFields.InUpperLayer = rv
		return nil
	case "open.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.Inode"}
		}
		ev.Open.File.FileFields.Inode = uint64(rv)
		return nil
	case "open.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.Mode"}
		}
		ev.Open.File.FileFields.Mode = uint16(rv)
		return nil
	case "open.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.MTime"}
		}
		ev.Open.File.FileFields.MTime = uint64(rv)
		return nil
	case "open.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.MountID"}
		}
		ev.Open.File.FileFields.MountID = uint32(rv)
		return nil
	case "open.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.BasenameStr"}
		}
		ev.Open.File.BasenameStr = rv
		return nil
	case "open.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "open.file.name.length"}
	case "open.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.PathnameStr"}
		}
		ev.Open.File.PathnameStr = rv
		return nil
	case "open.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "open.file.path.length"}
	case "open.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.Mode"}
		}
		ev.Open.File.FileFields.Mode = uint16(rv)
		return nil
	case "open.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.UID"}
		}
		ev.Open.File.FileFields.UID = uint32(rv)
		return nil
	case "open.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.File.FileFields.User"}
		}
		ev.Open.File.FileFields.User = rv
		return nil
	case "open.flags":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.Flags"}
		}
		ev.Open.Flags = uint32(rv)
		return nil
	case "open.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Open.SyscallEvent.Retval"}
		}
		ev.Open.SyscallEvent.Retval = int64(rv)
		return nil
	case "process.ancestors.args":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Args"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Args = rv
		return nil
	case "process.ancestors.args_flags":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Argv = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Argv, rv)
		case []string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Argv = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Argv"}
		}
		return nil
	case "process.ancestors.args_options":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Argv = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Argv, rv)
		case []string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Argv = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Argv"}
		}
		return nil
	case "process.ancestors.args_truncated":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.ArgsTruncated"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.ArgsTruncated = rv
		return nil
	case "process.ancestors.argv":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Argv = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Argv, rv)
		case []string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Argv = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Argv"}
		}
		return nil
	case "process.ancestors.argv0":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Argv0"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Argv0 = rv
		return nil
	case "process.ancestors.cap_effective":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.CapEffective"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.CapEffective = uint64(rv)
		return nil
	case "process.ancestors.cap_permitted":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.CapPermitted"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.CapPermitted = uint64(rv)
		return nil
	case "process.ancestors.comm":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Comm"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Comm = rv
		return nil
	case "process.ancestors.container.id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.ContainerID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.ContainerID = rv
		return nil
	case "process.ancestors.cookie":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Cookie"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Cookie = uint32(rv)
		return nil
	case "process.ancestors.created_at":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.CreatedAt"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.CreatedAt = uint64(rv)
		return nil
	case "process.ancestors.egid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.EGID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.EGID = uint32(rv)
		return nil
	case "process.ancestors.egroup":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.EGroup"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.EGroup = rv
		return nil
	case "process.ancestors.envp":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Envp = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Envp, rv)
		case []string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Envp = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Envp"}
		}
		return nil
	case "process.ancestors.envs":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Envs = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Envs, rv)
		case []string:
			ev.ProcessContext.Ancestor.ProcessContext.Process.Envs = append(ev.ProcessContext.Ancestor.ProcessContext.Process.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Envs"}
		}
		return nil
	case "process.ancestors.envs_truncated":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.EnvsTruncated"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.EnvsTruncated = rv
		return nil
	case "process.ancestors.euid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.EUID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.EUID = uint32(rv)
		return nil
	case "process.ancestors.euser":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.EUser"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.EUser = rv
		return nil
	case "process.ancestors.file.change_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.CTime"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "process.ancestors.file.filesystem":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.Filesystem"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.Filesystem = rv
		return nil
	case "process.ancestors.file.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.GID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "process.ancestors.file.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Group"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Group = rv
		return nil
	case "process.ancestors.file.in_upper_layer":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.InUpperLayer"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "process.ancestors.file.inode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Inode"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "process.ancestors.file.mode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.ancestors.file.modification_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.MTime"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "process.ancestors.file.mount_id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.MountID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "process.ancestors.file.name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.BasenameStr"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.BasenameStr = rv
		return nil
	case "process.ancestors.file.name.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.ancestors.file.name.length"}
	case "process.ancestors.file.path":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.PathnameStr"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.PathnameStr = rv
		return nil
	case "process.ancestors.file.path.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.ancestors.file.path.length"}
	case "process.ancestors.file.rights":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.ancestors.file.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.UID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "process.ancestors.file.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.User"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.FileEvent.FileFields.User = rv
		return nil
	case "process.ancestors.fsgid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSGID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSGID = uint32(rv)
		return nil
	case "process.ancestors.fsgroup":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSGroup"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSGroup = rv
		return nil
	case "process.ancestors.fsuid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSUID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSUID = uint32(rv)
		return nil
	case "process.ancestors.fsuser":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSUser"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.FSUser = rv
		return nil
	case "process.ancestors.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.GID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.GID = uint32(rv)
		return nil
	case "process.ancestors.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.Group"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.Group = rv
		return nil
	case "process.ancestors.interpreter.file.change_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "process.ancestors.interpreter.file.filesystem":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "process.ancestors.interpreter.file.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "process.ancestors.interpreter.file.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "process.ancestors.interpreter.file.in_upper_layer":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "process.ancestors.interpreter.file.inode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "process.ancestors.interpreter.file.mode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.ancestors.interpreter.file.modification_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "process.ancestors.interpreter.file.mount_id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "process.ancestors.interpreter.file.name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "process.ancestors.interpreter.file.name.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.ancestors.interpreter.file.name.length"}
	case "process.ancestors.interpreter.file.path":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "process.ancestors.interpreter.file.path.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.ancestors.interpreter.file.path.length"}
	case "process.ancestors.interpreter.file.rights":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.ancestors.interpreter.file.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "process.ancestors.interpreter.file.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "process.ancestors.is_kworker":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.PIDContext.IsKworker"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.PIDContext.IsKworker = rv
		return nil
	case "process.ancestors.is_thread":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.IsThread"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.IsThread = rv
		return nil
	case "process.ancestors.pid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.PIDContext.Pid"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.PIDContext.Pid = uint32(rv)
		return nil
	case "process.ancestors.ppid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.PPid"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.PPid = uint32(rv)
		return nil
	case "process.ancestors.tid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.PIDContext.Tid"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.PIDContext.Tid = uint32(rv)
		return nil
	case "process.ancestors.tty_name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.TTYName"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.TTYName = rv
		return nil
	case "process.ancestors.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.UID"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.UID = uint32(rv)
		return nil
	case "process.ancestors.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Ancestor == nil {
			ev.ProcessContext.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Ancestor.ProcessContext.Process.Credentials.User"}
		}
		ev.ProcessContext.Ancestor.ProcessContext.Process.Credentials.User = rv
		return nil
	case "process.args":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Args"}
		}
		ev.ProcessContext.Process.Args = rv
		return nil
	case "process.args_flags":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Process.Argv = append(ev.ProcessContext.Process.Argv, rv)
		case []string:
			ev.ProcessContext.Process.Argv = append(ev.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Argv"}
		}
		return nil
	case "process.args_options":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Process.Argv = append(ev.ProcessContext.Process.Argv, rv)
		case []string:
			ev.ProcessContext.Process.Argv = append(ev.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Argv"}
		}
		return nil
	case "process.args_truncated":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.ArgsTruncated"}
		}
		ev.ProcessContext.Process.ArgsTruncated = rv
		return nil
	case "process.argv":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Process.Argv = append(ev.ProcessContext.Process.Argv, rv)
		case []string:
			ev.ProcessContext.Process.Argv = append(ev.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Argv"}
		}
		return nil
	case "process.argv0":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Argv0"}
		}
		ev.ProcessContext.Process.Argv0 = rv
		return nil
	case "process.cap_effective":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.CapEffective"}
		}
		ev.ProcessContext.Process.Credentials.CapEffective = uint64(rv)
		return nil
	case "process.cap_permitted":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.CapPermitted"}
		}
		ev.ProcessContext.Process.Credentials.CapPermitted = uint64(rv)
		return nil
	case "process.comm":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Comm"}
		}
		ev.ProcessContext.Process.Comm = rv
		return nil
	case "process.container.id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.ContainerID"}
		}
		ev.ProcessContext.Process.ContainerID = rv
		return nil
	case "process.cookie":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Cookie"}
		}
		ev.ProcessContext.Process.Cookie = uint32(rv)
		return nil
	case "process.created_at":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.CreatedAt"}
		}
		ev.ProcessContext.Process.CreatedAt = uint64(rv)
		return nil
	case "process.egid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.EGID"}
		}
		ev.ProcessContext.Process.Credentials.EGID = uint32(rv)
		return nil
	case "process.egroup":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.EGroup"}
		}
		ev.ProcessContext.Process.Credentials.EGroup = rv
		return nil
	case "process.envp":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Process.Envp = append(ev.ProcessContext.Process.Envp, rv)
		case []string:
			ev.ProcessContext.Process.Envp = append(ev.ProcessContext.Process.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Envp"}
		}
		return nil
	case "process.envs":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Process.Envs = append(ev.ProcessContext.Process.Envs, rv)
		case []string:
			ev.ProcessContext.Process.Envs = append(ev.ProcessContext.Process.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Envs"}
		}
		return nil
	case "process.envs_truncated":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.EnvsTruncated"}
		}
		ev.ProcessContext.Process.EnvsTruncated = rv
		return nil
	case "process.euid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.EUID"}
		}
		ev.ProcessContext.Process.Credentials.EUID = uint32(rv)
		return nil
	case "process.euser":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.EUser"}
		}
		ev.ProcessContext.Process.Credentials.EUser = rv
		return nil
	case "process.file.change_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.CTime"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "process.file.filesystem":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.Filesystem"}
		}
		ev.ProcessContext.Process.FileEvent.Filesystem = rv
		return nil
	case "process.file.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.GID"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "process.file.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.Group"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.Group = rv
		return nil
	case "process.file.in_upper_layer":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.InUpperLayer"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "process.file.inode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.Inode"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "process.file.mode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.file.modification_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.MTime"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "process.file.mount_id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.MountID"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "process.file.name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.BasenameStr"}
		}
		ev.ProcessContext.Process.FileEvent.BasenameStr = rv
		return nil
	case "process.file.name.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.file.name.length"}
	case "process.file.path":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.PathnameStr"}
		}
		ev.ProcessContext.Process.FileEvent.PathnameStr = rv
		return nil
	case "process.file.path.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.file.path.length"}
	case "process.file.rights":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.file.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.UID"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "process.file.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.FileEvent.FileFields.User"}
		}
		ev.ProcessContext.Process.FileEvent.FileFields.User = rv
		return nil
	case "process.fsgid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.FSGID"}
		}
		ev.ProcessContext.Process.Credentials.FSGID = uint32(rv)
		return nil
	case "process.fsgroup":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.FSGroup"}
		}
		ev.ProcessContext.Process.Credentials.FSGroup = rv
		return nil
	case "process.fsuid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.FSUID"}
		}
		ev.ProcessContext.Process.Credentials.FSUID = uint32(rv)
		return nil
	case "process.fsuser":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.FSUser"}
		}
		ev.ProcessContext.Process.Credentials.FSUser = rv
		return nil
	case "process.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.GID"}
		}
		ev.ProcessContext.Process.Credentials.GID = uint32(rv)
		return nil
	case "process.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.Group"}
		}
		ev.ProcessContext.Process.Credentials.Group = rv
		return nil
	case "process.interpreter.file.change_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "process.interpreter.file.filesystem":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "process.interpreter.file.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "process.interpreter.file.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "process.interpreter.file.in_upper_layer":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "process.interpreter.file.inode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "process.interpreter.file.mode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.interpreter.file.modification_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "process.interpreter.file.mount_id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "process.interpreter.file.name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "process.interpreter.file.name.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.interpreter.file.name.length"}
	case "process.interpreter.file.path":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "process.interpreter.file.path.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.interpreter.file.path.length"}
	case "process.interpreter.file.rights":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.interpreter.file.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "process.interpreter.file.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "process.is_kworker":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.PIDContext.IsKworker"}
		}
		ev.ProcessContext.Process.PIDContext.IsKworker = rv
		return nil
	case "process.is_thread":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.IsThread"}
		}
		ev.ProcessContext.Process.IsThread = rv
		return nil
	case "process.parent.args":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Args"}
		}
		ev.ProcessContext.Parent.Args = rv
		return nil
	case "process.parent.args_flags":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Parent.Argv = append(ev.ProcessContext.Parent.Argv, rv)
		case []string:
			ev.ProcessContext.Parent.Argv = append(ev.ProcessContext.Parent.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Argv"}
		}
		return nil
	case "process.parent.args_options":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Parent.Argv = append(ev.ProcessContext.Parent.Argv, rv)
		case []string:
			ev.ProcessContext.Parent.Argv = append(ev.ProcessContext.Parent.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Argv"}
		}
		return nil
	case "process.parent.args_truncated":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.ArgsTruncated"}
		}
		ev.ProcessContext.Parent.ArgsTruncated = rv
		return nil
	case "process.parent.argv":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Parent.Argv = append(ev.ProcessContext.Parent.Argv, rv)
		case []string:
			ev.ProcessContext.Parent.Argv = append(ev.ProcessContext.Parent.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Argv"}
		}
		return nil
	case "process.parent.argv0":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Argv0"}
		}
		ev.ProcessContext.Parent.Argv0 = rv
		return nil
	case "process.parent.cap_effective":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.CapEffective"}
		}
		ev.ProcessContext.Parent.Credentials.CapEffective = uint64(rv)
		return nil
	case "process.parent.cap_permitted":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.CapPermitted"}
		}
		ev.ProcessContext.Parent.Credentials.CapPermitted = uint64(rv)
		return nil
	case "process.parent.comm":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Comm"}
		}
		ev.ProcessContext.Parent.Comm = rv
		return nil
	case "process.parent.container.id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.ContainerID"}
		}
		ev.ProcessContext.Parent.ContainerID = rv
		return nil
	case "process.parent.cookie":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Cookie"}
		}
		ev.ProcessContext.Parent.Cookie = uint32(rv)
		return nil
	case "process.parent.created_at":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.CreatedAt"}
		}
		ev.ProcessContext.Parent.CreatedAt = uint64(rv)
		return nil
	case "process.parent.egid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.EGID"}
		}
		ev.ProcessContext.Parent.Credentials.EGID = uint32(rv)
		return nil
	case "process.parent.egroup":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.EGroup"}
		}
		ev.ProcessContext.Parent.Credentials.EGroup = rv
		return nil
	case "process.parent.envp":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Parent.Envp = append(ev.ProcessContext.Parent.Envp, rv)
		case []string:
			ev.ProcessContext.Parent.Envp = append(ev.ProcessContext.Parent.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Envp"}
		}
		return nil
	case "process.parent.envs":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.ProcessContext.Parent.Envs = append(ev.ProcessContext.Parent.Envs, rv)
		case []string:
			ev.ProcessContext.Parent.Envs = append(ev.ProcessContext.Parent.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Envs"}
		}
		return nil
	case "process.parent.envs_truncated":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.EnvsTruncated"}
		}
		ev.ProcessContext.Parent.EnvsTruncated = rv
		return nil
	case "process.parent.euid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.EUID"}
		}
		ev.ProcessContext.Parent.Credentials.EUID = uint32(rv)
		return nil
	case "process.parent.euser":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.EUser"}
		}
		ev.ProcessContext.Parent.Credentials.EUser = rv
		return nil
	case "process.parent.file.change_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.CTime"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "process.parent.file.filesystem":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.Filesystem"}
		}
		ev.ProcessContext.Parent.FileEvent.Filesystem = rv
		return nil
	case "process.parent.file.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.GID"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "process.parent.file.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.Group"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.Group = rv
		return nil
	case "process.parent.file.in_upper_layer":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.InUpperLayer"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "process.parent.file.inode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.Inode"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "process.parent.file.mode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.parent.file.modification_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.MTime"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "process.parent.file.mount_id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.MountID"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "process.parent.file.name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.BasenameStr"}
		}
		ev.ProcessContext.Parent.FileEvent.BasenameStr = rv
		return nil
	case "process.parent.file.name.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.parent.file.name.length"}
	case "process.parent.file.path":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.PathnameStr"}
		}
		ev.ProcessContext.Parent.FileEvent.PathnameStr = rv
		return nil
	case "process.parent.file.path.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.parent.file.path.length"}
	case "process.parent.file.rights":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.parent.file.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.UID"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "process.parent.file.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.FileEvent.FileFields.User"}
		}
		ev.ProcessContext.Parent.FileEvent.FileFields.User = rv
		return nil
	case "process.parent.fsgid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.FSGID"}
		}
		ev.ProcessContext.Parent.Credentials.FSGID = uint32(rv)
		return nil
	case "process.parent.fsgroup":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.FSGroup"}
		}
		ev.ProcessContext.Parent.Credentials.FSGroup = rv
		return nil
	case "process.parent.fsuid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.FSUID"}
		}
		ev.ProcessContext.Parent.Credentials.FSUID = uint32(rv)
		return nil
	case "process.parent.fsuser":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.FSUser"}
		}
		ev.ProcessContext.Parent.Credentials.FSUser = rv
		return nil
	case "process.parent.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.GID"}
		}
		ev.ProcessContext.Parent.Credentials.GID = uint32(rv)
		return nil
	case "process.parent.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.Group"}
		}
		ev.ProcessContext.Parent.Credentials.Group = rv
		return nil
	case "process.parent.interpreter.file.change_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "process.parent.interpreter.file.filesystem":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "process.parent.interpreter.file.gid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "process.parent.interpreter.file.group":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "process.parent.interpreter.file.in_upper_layer":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "process.parent.interpreter.file.inode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "process.parent.interpreter.file.mode":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.parent.interpreter.file.modification_time":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "process.parent.interpreter.file.mount_id":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "process.parent.interpreter.file.name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "process.parent.interpreter.file.name.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.parent.interpreter.file.name.length"}
	case "process.parent.interpreter.file.path":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "process.parent.interpreter.file.path.length":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "process.parent.interpreter.file.path.length"}
	case "process.parent.interpreter.file.rights":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "process.parent.interpreter.file.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "process.parent.interpreter.file.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.ProcessContext.Parent.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "process.parent.is_kworker":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.PIDContext.IsKworker"}
		}
		ev.ProcessContext.Parent.PIDContext.IsKworker = rv
		return nil
	case "process.parent.is_thread":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.IsThread"}
		}
		ev.ProcessContext.Parent.IsThread = rv
		return nil
	case "process.parent.pid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.PIDContext.Pid"}
		}
		ev.ProcessContext.Parent.PIDContext.Pid = uint32(rv)
		return nil
	case "process.parent.ppid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.PPid"}
		}
		ev.ProcessContext.Parent.PPid = uint32(rv)
		return nil
	case "process.parent.tid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.PIDContext.Tid"}
		}
		ev.ProcessContext.Parent.PIDContext.Tid = uint32(rv)
		return nil
	case "process.parent.tty_name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.TTYName"}
		}
		ev.ProcessContext.Parent.TTYName = rv
		return nil
	case "process.parent.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.UID"}
		}
		ev.ProcessContext.Parent.Credentials.UID = uint32(rv)
		return nil
	case "process.parent.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		if ev.ProcessContext.Parent == nil {
			ev.ProcessContext.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Parent.Credentials.User"}
		}
		ev.ProcessContext.Parent.Credentials.User = rv
		return nil
	case "process.pid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.PIDContext.Pid"}
		}
		ev.ProcessContext.Process.PIDContext.Pid = uint32(rv)
		return nil
	case "process.ppid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.PPid"}
		}
		ev.ProcessContext.Process.PPid = uint32(rv)
		return nil
	case "process.tid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.PIDContext.Tid"}
		}
		ev.ProcessContext.Process.PIDContext.Tid = uint32(rv)
		return nil
	case "process.tty_name":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.TTYName"}
		}
		ev.ProcessContext.Process.TTYName = rv
		return nil
	case "process.uid":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.UID"}
		}
		ev.ProcessContext.Process.Credentials.UID = uint32(rv)
		return nil
	case "process.user":
		if ev.ProcessContext == nil {
			ev.ProcessContext = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "ProcessContext.Process.Credentials.User"}
		}
		ev.ProcessContext.Process.Credentials.User = rv
		return nil
	case "removexattr.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.CTime"}
		}
		ev.RemoveXAttr.File.FileFields.CTime = uint64(rv)
		return nil
	case "removexattr.file.destination.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.Name"}
		}
		ev.RemoveXAttr.Name = rv
		return nil
	case "removexattr.file.destination.namespace":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.Namespace"}
		}
		ev.RemoveXAttr.Namespace = rv
		return nil
	case "removexattr.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.Filesystem"}
		}
		ev.RemoveXAttr.File.Filesystem = rv
		return nil
	case "removexattr.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.GID"}
		}
		ev.RemoveXAttr.File.FileFields.GID = uint32(rv)
		return nil
	case "removexattr.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.Group"}
		}
		ev.RemoveXAttr.File.FileFields.Group = rv
		return nil
	case "removexattr.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.InUpperLayer"}
		}
		ev.RemoveXAttr.File.FileFields.InUpperLayer = rv
		return nil
	case "removexattr.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.Inode"}
		}
		ev.RemoveXAttr.File.FileFields.Inode = uint64(rv)
		return nil
	case "removexattr.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.Mode"}
		}
		ev.RemoveXAttr.File.FileFields.Mode = uint16(rv)
		return nil
	case "removexattr.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.MTime"}
		}
		ev.RemoveXAttr.File.FileFields.MTime = uint64(rv)
		return nil
	case "removexattr.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.MountID"}
		}
		ev.RemoveXAttr.File.FileFields.MountID = uint32(rv)
		return nil
	case "removexattr.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.BasenameStr"}
		}
		ev.RemoveXAttr.File.BasenameStr = rv
		return nil
	case "removexattr.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "removexattr.file.name.length"}
	case "removexattr.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.PathnameStr"}
		}
		ev.RemoveXAttr.File.PathnameStr = rv
		return nil
	case "removexattr.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "removexattr.file.path.length"}
	case "removexattr.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.Mode"}
		}
		ev.RemoveXAttr.File.FileFields.Mode = uint16(rv)
		return nil
	case "removexattr.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.UID"}
		}
		ev.RemoveXAttr.File.FileFields.UID = uint32(rv)
		return nil
	case "removexattr.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.File.FileFields.User"}
		}
		ev.RemoveXAttr.File.FileFields.User = rv
		return nil
	case "removexattr.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "RemoveXAttr.SyscallEvent.Retval"}
		}
		ev.RemoveXAttr.SyscallEvent.Retval = int64(rv)
		return nil
	case "rename.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.CTime"}
		}
		ev.Rename.Old.FileFields.CTime = uint64(rv)
		return nil
	case "rename.file.destination.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.CTime"}
		}
		ev.Rename.New.FileFields.CTime = uint64(rv)
		return nil
	case "rename.file.destination.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.Filesystem"}
		}
		ev.Rename.New.Filesystem = rv
		return nil
	case "rename.file.destination.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.GID"}
		}
		ev.Rename.New.FileFields.GID = uint32(rv)
		return nil
	case "rename.file.destination.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.Group"}
		}
		ev.Rename.New.FileFields.Group = rv
		return nil
	case "rename.file.destination.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.InUpperLayer"}
		}
		ev.Rename.New.FileFields.InUpperLayer = rv
		return nil
	case "rename.file.destination.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.Inode"}
		}
		ev.Rename.New.FileFields.Inode = uint64(rv)
		return nil
	case "rename.file.destination.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.Mode"}
		}
		ev.Rename.New.FileFields.Mode = uint16(rv)
		return nil
	case "rename.file.destination.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.MTime"}
		}
		ev.Rename.New.FileFields.MTime = uint64(rv)
		return nil
	case "rename.file.destination.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.MountID"}
		}
		ev.Rename.New.FileFields.MountID = uint32(rv)
		return nil
	case "rename.file.destination.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.BasenameStr"}
		}
		ev.Rename.New.BasenameStr = rv
		return nil
	case "rename.file.destination.name.length":
		return &eval.ErrFieldReadOnly{Field: "rename.file.destination.name.length"}
	case "rename.file.destination.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.PathnameStr"}
		}
		ev.Rename.New.PathnameStr = rv
		return nil
	case "rename.file.destination.path.length":
		return &eval.ErrFieldReadOnly{Field: "rename.file.destination.path.length"}
	case "rename.file.destination.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.Mode"}
		}
		ev.Rename.New.FileFields.Mode = uint16(rv)
		return nil
	case "rename.file.destination.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.UID"}
		}
		ev.Rename.New.FileFields.UID = uint32(rv)
		return nil
	case "rename.file.destination.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.New.FileFields.User"}
		}
		ev.Rename.New.FileFields.User = rv
		return nil
	case "rename.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.Filesystem"}
		}
		ev.Rename.Old.Filesystem = rv
		return nil
	case "rename.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.GID"}
		}
		ev.Rename.Old.FileFields.GID = uint32(rv)
		return nil
	case "rename.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.Group"}
		}
		ev.Rename.Old.FileFields.Group = rv
		return nil
	case "rename.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.InUpperLayer"}
		}
		ev.Rename.Old.FileFields.InUpperLayer = rv
		return nil
	case "rename.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.Inode"}
		}
		ev.Rename.Old.FileFields.Inode = uint64(rv)
		return nil
	case "rename.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.Mode"}
		}
		ev.Rename.Old.FileFields.Mode = uint16(rv)
		return nil
	case "rename.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.MTime"}
		}
		ev.Rename.Old.FileFields.MTime = uint64(rv)
		return nil
	case "rename.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.MountID"}
		}
		ev.Rename.Old.FileFields.MountID = uint32(rv)
		return nil
	case "rename.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.BasenameStr"}
		}
		ev.Rename.Old.BasenameStr = rv
		return nil
	case "rename.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "rename.file.name.length"}
	case "rename.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.PathnameStr"}
		}
		ev.Rename.Old.PathnameStr = rv
		return nil
	case "rename.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "rename.file.path.length"}
	case "rename.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.Mode"}
		}
		ev.Rename.Old.FileFields.Mode = uint16(rv)
		return nil
	case "rename.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.UID"}
		}
		ev.Rename.Old.FileFields.UID = uint32(rv)
		return nil
	case "rename.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.Old.FileFields.User"}
		}
		ev.Rename.Old.FileFields.User = rv
		return nil
	case "rename.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rename.SyscallEvent.Retval"}
		}
		ev.Rename.SyscallEvent.Retval = int64(rv)
		return nil
	case "rmdir.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.CTime"}
		}
		ev.Rmdir.File.FileFields.CTime = uint64(rv)
		return nil
	case "rmdir.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.Filesystem"}
		}
		ev.Rmdir.File.Filesystem = rv
		return nil
	case "rmdir.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.GID"}
		}
		ev.Rmdir.File.FileFields.GID = uint32(rv)
		return nil
	case "rmdir.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.Group"}
		}
		ev.Rmdir.File.FileFields.Group = rv
		return nil
	case "rmdir.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.InUpperLayer"}
		}
		ev.Rmdir.File.FileFields.InUpperLayer = rv
		return nil
	case "rmdir.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.Inode"}
		}
		ev.Rmdir.File.FileFields.Inode = uint64(rv)
		return nil
	case "rmdir.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.Mode"}
		}
		ev.Rmdir.File.FileFields.Mode = uint16(rv)
		return nil
	case "rmdir.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.MTime"}
		}
		ev.Rmdir.File.FileFields.MTime = uint64(rv)
		return nil
	case "rmdir.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.MountID"}
		}
		ev.Rmdir.File.FileFields.MountID = uint32(rv)
		return nil
	case "rmdir.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.BasenameStr"}
		}
		ev.Rmdir.File.BasenameStr = rv
		return nil
	case "rmdir.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "rmdir.file.name.length"}
	case "rmdir.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.PathnameStr"}
		}
		ev.Rmdir.File.PathnameStr = rv
		return nil
	case "rmdir.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "rmdir.file.path.length"}
	case "rmdir.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.Mode"}
		}
		ev.Rmdir.File.FileFields.Mode = uint16(rv)
		return nil
	case "rmdir.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.UID"}
		}
		ev.Rmdir.File.FileFields.UID = uint32(rv)
		return nil
	case "rmdir.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.File.FileFields.User"}
		}
		ev.Rmdir.File.FileFields.User = rv
		return nil
	case "rmdir.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Rmdir.SyscallEvent.Retval"}
		}
		ev.Rmdir.SyscallEvent.Retval = int64(rv)
		return nil
	case "setgid.egid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetGID.EGID"}
		}
		ev.SetGID.EGID = uint32(rv)
		return nil
	case "setgid.egroup":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetGID.EGroup"}
		}
		ev.SetGID.EGroup = rv
		return nil
	case "setgid.fsgid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetGID.FSGID"}
		}
		ev.SetGID.FSGID = uint32(rv)
		return nil
	case "setgid.fsgroup":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetGID.FSGroup"}
		}
		ev.SetGID.FSGroup = rv
		return nil
	case "setgid.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetGID.GID"}
		}
		ev.SetGID.GID = uint32(rv)
		return nil
	case "setgid.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetGID.Group"}
		}
		ev.SetGID.Group = rv
		return nil
	case "setuid.euid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetUID.EUID"}
		}
		ev.SetUID.EUID = uint32(rv)
		return nil
	case "setuid.euser":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetUID.EUser"}
		}
		ev.SetUID.EUser = rv
		return nil
	case "setuid.fsuid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetUID.FSUID"}
		}
		ev.SetUID.FSUID = uint32(rv)
		return nil
	case "setuid.fsuser":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetUID.FSUser"}
		}
		ev.SetUID.FSUser = rv
		return nil
	case "setuid.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetUID.UID"}
		}
		ev.SetUID.UID = uint32(rv)
		return nil
	case "setuid.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetUID.User"}
		}
		ev.SetUID.User = rv
		return nil
	case "setxattr.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.CTime"}
		}
		ev.SetXAttr.File.FileFields.CTime = uint64(rv)
		return nil
	case "setxattr.file.destination.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.Name"}
		}
		ev.SetXAttr.Name = rv
		return nil
	case "setxattr.file.destination.namespace":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.Namespace"}
		}
		ev.SetXAttr.Namespace = rv
		return nil
	case "setxattr.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.Filesystem"}
		}
		ev.SetXAttr.File.Filesystem = rv
		return nil
	case "setxattr.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.GID"}
		}
		ev.SetXAttr.File.FileFields.GID = uint32(rv)
		return nil
	case "setxattr.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.Group"}
		}
		ev.SetXAttr.File.FileFields.Group = rv
		return nil
	case "setxattr.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.InUpperLayer"}
		}
		ev.SetXAttr.File.FileFields.InUpperLayer = rv
		return nil
	case "setxattr.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.Inode"}
		}
		ev.SetXAttr.File.FileFields.Inode = uint64(rv)
		return nil
	case "setxattr.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.Mode"}
		}
		ev.SetXAttr.File.FileFields.Mode = uint16(rv)
		return nil
	case "setxattr.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.MTime"}
		}
		ev.SetXAttr.File.FileFields.MTime = uint64(rv)
		return nil
	case "setxattr.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.MountID"}
		}
		ev.SetXAttr.File.FileFields.MountID = uint32(rv)
		return nil
	case "setxattr.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.BasenameStr"}
		}
		ev.SetXAttr.File.BasenameStr = rv
		return nil
	case "setxattr.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "setxattr.file.name.length"}
	case "setxattr.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.PathnameStr"}
		}
		ev.SetXAttr.File.PathnameStr = rv
		return nil
	case "setxattr.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "setxattr.file.path.length"}
	case "setxattr.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.Mode"}
		}
		ev.SetXAttr.File.FileFields.Mode = uint16(rv)
		return nil
	case "setxattr.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.UID"}
		}
		ev.SetXAttr.File.FileFields.UID = uint32(rv)
		return nil
	case "setxattr.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.File.FileFields.User"}
		}
		ev.SetXAttr.File.FileFields.User = rv
		return nil
	case "setxattr.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "SetXAttr.SyscallEvent.Retval"}
		}
		ev.SetXAttr.SyscallEvent.Retval = int64(rv)
		return nil
	case "signal.pid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.PID"}
		}
		ev.Signal.PID = uint32(rv)
		return nil
	case "signal.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.SyscallEvent.Retval"}
		}
		ev.Signal.SyscallEvent.Retval = int64(rv)
		return nil
	case "signal.target.ancestors.args":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Args"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Args = rv
		return nil
	case "signal.target.ancestors.args_flags":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Argv = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Argv, rv)
		case []string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Argv = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Argv"}
		}
		return nil
	case "signal.target.ancestors.args_options":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Argv = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Argv, rv)
		case []string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Argv = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Argv"}
		}
		return nil
	case "signal.target.ancestors.args_truncated":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.ArgsTruncated"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.ArgsTruncated = rv
		return nil
	case "signal.target.ancestors.argv":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Argv = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Argv, rv)
		case []string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Argv = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Argv"}
		}
		return nil
	case "signal.target.ancestors.argv0":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Argv0"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Argv0 = rv
		return nil
	case "signal.target.ancestors.cap_effective":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.CapEffective"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.CapEffective = uint64(rv)
		return nil
	case "signal.target.ancestors.cap_permitted":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.CapPermitted"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.CapPermitted = uint64(rv)
		return nil
	case "signal.target.ancestors.comm":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Comm"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Comm = rv
		return nil
	case "signal.target.ancestors.container.id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.ContainerID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.ContainerID = rv
		return nil
	case "signal.target.ancestors.cookie":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Cookie"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Cookie = uint32(rv)
		return nil
	case "signal.target.ancestors.created_at":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.CreatedAt"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.CreatedAt = uint64(rv)
		return nil
	case "signal.target.ancestors.egid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.EGID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.EGID = uint32(rv)
		return nil
	case "signal.target.ancestors.egroup":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.EGroup"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.EGroup = rv
		return nil
	case "signal.target.ancestors.envp":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Envp = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Envp, rv)
		case []string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Envp = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Envp"}
		}
		return nil
	case "signal.target.ancestors.envs":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Envs = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Envs, rv)
		case []string:
			ev.Signal.Target.Ancestor.ProcessContext.Process.Envs = append(ev.Signal.Target.Ancestor.ProcessContext.Process.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Envs"}
		}
		return nil
	case "signal.target.ancestors.envs_truncated":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.EnvsTruncated"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.EnvsTruncated = rv
		return nil
	case "signal.target.ancestors.euid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.EUID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.EUID = uint32(rv)
		return nil
	case "signal.target.ancestors.euser":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.EUser"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.EUser = rv
		return nil
	case "signal.target.ancestors.file.change_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.CTime"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "signal.target.ancestors.file.filesystem":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.Filesystem"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.Filesystem = rv
		return nil
	case "signal.target.ancestors.file.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.GID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "signal.target.ancestors.file.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Group"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Group = rv
		return nil
	case "signal.target.ancestors.file.in_upper_layer":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "signal.target.ancestors.file.inode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Inode"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "signal.target.ancestors.file.mode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.ancestors.file.modification_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.MTime"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "signal.target.ancestors.file.mount_id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.MountID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "signal.target.ancestors.file.name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.BasenameStr"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.BasenameStr = rv
		return nil
	case "signal.target.ancestors.file.name.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.ancestors.file.name.length"}
	case "signal.target.ancestors.file.path":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.PathnameStr"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.PathnameStr = rv
		return nil
	case "signal.target.ancestors.file.path.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.ancestors.file.path.length"}
	case "signal.target.ancestors.file.rights":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.ancestors.file.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.UID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "signal.target.ancestors.file.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.User"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.FileEvent.FileFields.User = rv
		return nil
	case "signal.target.ancestors.fsgid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSGID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSGID = uint32(rv)
		return nil
	case "signal.target.ancestors.fsgroup":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSGroup"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSGroup = rv
		return nil
	case "signal.target.ancestors.fsuid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSUID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSUID = uint32(rv)
		return nil
	case "signal.target.ancestors.fsuser":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSUser"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.FSUser = rv
		return nil
	case "signal.target.ancestors.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.GID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.GID = uint32(rv)
		return nil
	case "signal.target.ancestors.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.Group"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.Group = rv
		return nil
	case "signal.target.ancestors.interpreter.file.change_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.filesystem":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "signal.target.ancestors.interpreter.file.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "signal.target.ancestors.interpreter.file.in_upper_layer":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "signal.target.ancestors.interpreter.file.inode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.mode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.modification_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.mount_id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "signal.target.ancestors.interpreter.file.name.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.ancestors.interpreter.file.name.length"}
	case "signal.target.ancestors.interpreter.file.path":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "signal.target.ancestors.interpreter.file.path.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.ancestors.interpreter.file.path.length"}
	case "signal.target.ancestors.interpreter.file.rights":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "signal.target.ancestors.interpreter.file.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "signal.target.ancestors.is_kworker":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.PIDContext.IsKworker"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.PIDContext.IsKworker = rv
		return nil
	case "signal.target.ancestors.is_thread":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.IsThread"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.IsThread = rv
		return nil
	case "signal.target.ancestors.pid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.PIDContext.Pid"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.PIDContext.Pid = uint32(rv)
		return nil
	case "signal.target.ancestors.ppid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.PPid"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.PPid = uint32(rv)
		return nil
	case "signal.target.ancestors.tid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.PIDContext.Tid"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.PIDContext.Tid = uint32(rv)
		return nil
	case "signal.target.ancestors.tty_name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.TTYName"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.TTYName = rv
		return nil
	case "signal.target.ancestors.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.UID"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.UID = uint32(rv)
		return nil
	case "signal.target.ancestors.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Ancestor == nil {
			ev.Signal.Target.Ancestor = &ProcessCacheEntry{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Ancestor.ProcessContext.Process.Credentials.User"}
		}
		ev.Signal.Target.Ancestor.ProcessContext.Process.Credentials.User = rv
		return nil
	case "signal.target.args":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Args"}
		}
		ev.Signal.Target.Process.Args = rv
		return nil
	case "signal.target.args_flags":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Process.Argv = append(ev.Signal.Target.Process.Argv, rv)
		case []string:
			ev.Signal.Target.Process.Argv = append(ev.Signal.Target.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Argv"}
		}
		return nil
	case "signal.target.args_options":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Process.Argv = append(ev.Signal.Target.Process.Argv, rv)
		case []string:
			ev.Signal.Target.Process.Argv = append(ev.Signal.Target.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Argv"}
		}
		return nil
	case "signal.target.args_truncated":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.ArgsTruncated"}
		}
		ev.Signal.Target.Process.ArgsTruncated = rv
		return nil
	case "signal.target.argv":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Process.Argv = append(ev.Signal.Target.Process.Argv, rv)
		case []string:
			ev.Signal.Target.Process.Argv = append(ev.Signal.Target.Process.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Argv"}
		}
		return nil
	case "signal.target.argv0":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Argv0"}
		}
		ev.Signal.Target.Process.Argv0 = rv
		return nil
	case "signal.target.cap_effective":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.CapEffective"}
		}
		ev.Signal.Target.Process.Credentials.CapEffective = uint64(rv)
		return nil
	case "signal.target.cap_permitted":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.CapPermitted"}
		}
		ev.Signal.Target.Process.Credentials.CapPermitted = uint64(rv)
		return nil
	case "signal.target.comm":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Comm"}
		}
		ev.Signal.Target.Process.Comm = rv
		return nil
	case "signal.target.container.id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.ContainerID"}
		}
		ev.Signal.Target.Process.ContainerID = rv
		return nil
	case "signal.target.cookie":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Cookie"}
		}
		ev.Signal.Target.Process.Cookie = uint32(rv)
		return nil
	case "signal.target.created_at":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.CreatedAt"}
		}
		ev.Signal.Target.Process.CreatedAt = uint64(rv)
		return nil
	case "signal.target.egid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.EGID"}
		}
		ev.Signal.Target.Process.Credentials.EGID = uint32(rv)
		return nil
	case "signal.target.egroup":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.EGroup"}
		}
		ev.Signal.Target.Process.Credentials.EGroup = rv
		return nil
	case "signal.target.envp":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Process.Envp = append(ev.Signal.Target.Process.Envp, rv)
		case []string:
			ev.Signal.Target.Process.Envp = append(ev.Signal.Target.Process.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Envp"}
		}
		return nil
	case "signal.target.envs":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Process.Envs = append(ev.Signal.Target.Process.Envs, rv)
		case []string:
			ev.Signal.Target.Process.Envs = append(ev.Signal.Target.Process.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Envs"}
		}
		return nil
	case "signal.target.envs_truncated":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.EnvsTruncated"}
		}
		ev.Signal.Target.Process.EnvsTruncated = rv
		return nil
	case "signal.target.euid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.EUID"}
		}
		ev.Signal.Target.Process.Credentials.EUID = uint32(rv)
		return nil
	case "signal.target.euser":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.EUser"}
		}
		ev.Signal.Target.Process.Credentials.EUser = rv
		return nil
	case "signal.target.file.change_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.CTime"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "signal.target.file.filesystem":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.Filesystem"}
		}
		ev.Signal.Target.Process.FileEvent.Filesystem = rv
		return nil
	case "signal.target.file.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.GID"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "signal.target.file.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.Group"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.Group = rv
		return nil
	case "signal.target.file.in_upper_layer":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "signal.target.file.inode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.Inode"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "signal.target.file.mode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.file.modification_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.MTime"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "signal.target.file.mount_id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.MountID"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "signal.target.file.name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.BasenameStr"}
		}
		ev.Signal.Target.Process.FileEvent.BasenameStr = rv
		return nil
	case "signal.target.file.name.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.file.name.length"}
	case "signal.target.file.path":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.PathnameStr"}
		}
		ev.Signal.Target.Process.FileEvent.PathnameStr = rv
		return nil
	case "signal.target.file.path.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.file.path.length"}
	case "signal.target.file.rights":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.file.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.UID"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "signal.target.file.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.FileEvent.FileFields.User"}
		}
		ev.Signal.Target.Process.FileEvent.FileFields.User = rv
		return nil
	case "signal.target.fsgid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.FSGID"}
		}
		ev.Signal.Target.Process.Credentials.FSGID = uint32(rv)
		return nil
	case "signal.target.fsgroup":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.FSGroup"}
		}
		ev.Signal.Target.Process.Credentials.FSGroup = rv
		return nil
	case "signal.target.fsuid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.FSUID"}
		}
		ev.Signal.Target.Process.Credentials.FSUID = uint32(rv)
		return nil
	case "signal.target.fsuser":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.FSUser"}
		}
		ev.Signal.Target.Process.Credentials.FSUser = rv
		return nil
	case "signal.target.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.GID"}
		}
		ev.Signal.Target.Process.Credentials.GID = uint32(rv)
		return nil
	case "signal.target.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.Group"}
		}
		ev.Signal.Target.Process.Credentials.Group = rv
		return nil
	case "signal.target.interpreter.file.change_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "signal.target.interpreter.file.filesystem":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "signal.target.interpreter.file.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "signal.target.interpreter.file.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "signal.target.interpreter.file.in_upper_layer":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "signal.target.interpreter.file.inode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "signal.target.interpreter.file.mode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.interpreter.file.modification_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "signal.target.interpreter.file.mount_id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "signal.target.interpreter.file.name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "signal.target.interpreter.file.name.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.interpreter.file.name.length"}
	case "signal.target.interpreter.file.path":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "signal.target.interpreter.file.path.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.interpreter.file.path.length"}
	case "signal.target.interpreter.file.rights":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.interpreter.file.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "signal.target.interpreter.file.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.Signal.Target.Process.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "signal.target.is_kworker":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.PIDContext.IsKworker"}
		}
		ev.Signal.Target.Process.PIDContext.IsKworker = rv
		return nil
	case "signal.target.is_thread":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.IsThread"}
		}
		ev.Signal.Target.Process.IsThread = rv
		return nil
	case "signal.target.parent.args":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Args"}
		}
		ev.Signal.Target.Parent.Args = rv
		return nil
	case "signal.target.parent.args_flags":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Parent.Argv = append(ev.Signal.Target.Parent.Argv, rv)
		case []string:
			ev.Signal.Target.Parent.Argv = append(ev.Signal.Target.Parent.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Argv"}
		}
		return nil
	case "signal.target.parent.args_options":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Parent.Argv = append(ev.Signal.Target.Parent.Argv, rv)
		case []string:
			ev.Signal.Target.Parent.Argv = append(ev.Signal.Target.Parent.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Argv"}
		}
		return nil
	case "signal.target.parent.args_truncated":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.ArgsTruncated"}
		}
		ev.Signal.Target.Parent.ArgsTruncated = rv
		return nil
	case "signal.target.parent.argv":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Parent.Argv = append(ev.Signal.Target.Parent.Argv, rv)
		case []string:
			ev.Signal.Target.Parent.Argv = append(ev.Signal.Target.Parent.Argv, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Argv"}
		}
		return nil
	case "signal.target.parent.argv0":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Argv0"}
		}
		ev.Signal.Target.Parent.Argv0 = rv
		return nil
	case "signal.target.parent.cap_effective":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.CapEffective"}
		}
		ev.Signal.Target.Parent.Credentials.CapEffective = uint64(rv)
		return nil
	case "signal.target.parent.cap_permitted":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.CapPermitted"}
		}
		ev.Signal.Target.Parent.Credentials.CapPermitted = uint64(rv)
		return nil
	case "signal.target.parent.comm":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Comm"}
		}
		ev.Signal.Target.Parent.Comm = rv
		return nil
	case "signal.target.parent.container.id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.ContainerID"}
		}
		ev.Signal.Target.Parent.ContainerID = rv
		return nil
	case "signal.target.parent.cookie":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Cookie"}
		}
		ev.Signal.Target.Parent.Cookie = uint32(rv)
		return nil
	case "signal.target.parent.created_at":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.CreatedAt"}
		}
		ev.Signal.Target.Parent.CreatedAt = uint64(rv)
		return nil
	case "signal.target.parent.egid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.EGID"}
		}
		ev.Signal.Target.Parent.Credentials.EGID = uint32(rv)
		return nil
	case "signal.target.parent.egroup":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.EGroup"}
		}
		ev.Signal.Target.Parent.Credentials.EGroup = rv
		return nil
	case "signal.target.parent.envp":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Parent.Envp = append(ev.Signal.Target.Parent.Envp, rv)
		case []string:
			ev.Signal.Target.Parent.Envp = append(ev.Signal.Target.Parent.Envp, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Envp"}
		}
		return nil
	case "signal.target.parent.envs":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		switch rv := value.(type) {
		case string:
			ev.Signal.Target.Parent.Envs = append(ev.Signal.Target.Parent.Envs, rv)
		case []string:
			ev.Signal.Target.Parent.Envs = append(ev.Signal.Target.Parent.Envs, rv...)
		default:
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Envs"}
		}
		return nil
	case "signal.target.parent.envs_truncated":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.EnvsTruncated"}
		}
		ev.Signal.Target.Parent.EnvsTruncated = rv
		return nil
	case "signal.target.parent.euid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.EUID"}
		}
		ev.Signal.Target.Parent.Credentials.EUID = uint32(rv)
		return nil
	case "signal.target.parent.euser":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.EUser"}
		}
		ev.Signal.Target.Parent.Credentials.EUser = rv
		return nil
	case "signal.target.parent.file.change_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.CTime"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "signal.target.parent.file.filesystem":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.Filesystem"}
		}
		ev.Signal.Target.Parent.FileEvent.Filesystem = rv
		return nil
	case "signal.target.parent.file.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.GID"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "signal.target.parent.file.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.Group"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.Group = rv
		return nil
	case "signal.target.parent.file.in_upper_layer":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "signal.target.parent.file.inode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.Inode"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "signal.target.parent.file.mode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.parent.file.modification_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.MTime"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "signal.target.parent.file.mount_id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.MountID"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "signal.target.parent.file.name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.BasenameStr"}
		}
		ev.Signal.Target.Parent.FileEvent.BasenameStr = rv
		return nil
	case "signal.target.parent.file.name.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.parent.file.name.length"}
	case "signal.target.parent.file.path":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.PathnameStr"}
		}
		ev.Signal.Target.Parent.FileEvent.PathnameStr = rv
		return nil
	case "signal.target.parent.file.path.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.parent.file.path.length"}
	case "signal.target.parent.file.rights":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.parent.file.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.UID"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "signal.target.parent.file.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.FileEvent.FileFields.User"}
		}
		ev.Signal.Target.Parent.FileEvent.FileFields.User = rv
		return nil
	case "signal.target.parent.fsgid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.FSGID"}
		}
		ev.Signal.Target.Parent.Credentials.FSGID = uint32(rv)
		return nil
	case "signal.target.parent.fsgroup":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.FSGroup"}
		}
		ev.Signal.Target.Parent.Credentials.FSGroup = rv
		return nil
	case "signal.target.parent.fsuid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.FSUID"}
		}
		ev.Signal.Target.Parent.Credentials.FSUID = uint32(rv)
		return nil
	case "signal.target.parent.fsuser":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.FSUser"}
		}
		ev.Signal.Target.Parent.Credentials.FSUser = rv
		return nil
	case "signal.target.parent.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.GID"}
		}
		ev.Signal.Target.Parent.Credentials.GID = uint32(rv)
		return nil
	case "signal.target.parent.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.Group"}
		}
		ev.Signal.Target.Parent.Credentials.Group = rv
		return nil
	case "signal.target.parent.interpreter.file.change_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.CTime"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.CTime = uint64(rv)
		return nil
	case "signal.target.parent.interpreter.file.filesystem":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.Filesystem"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.Filesystem = rv
		return nil
	case "signal.target.parent.interpreter.file.gid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.GID"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.GID = uint32(rv)
		return nil
	case "signal.target.parent.interpreter.file.group":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Group"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Group = rv
		return nil
	case "signal.target.parent.interpreter.file.in_upper_layer":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.InUpperLayer"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.InUpperLayer = rv
		return nil
	case "signal.target.parent.interpreter.file.inode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Inode"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Inode = uint64(rv)
		return nil
	case "signal.target.parent.interpreter.file.mode":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.parent.interpreter.file.modification_time":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MTime"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MTime = uint64(rv)
		return nil
	case "signal.target.parent.interpreter.file.mount_id":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MountID"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.MountID = uint32(rv)
		return nil
	case "signal.target.parent.interpreter.file.name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.BasenameStr"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.BasenameStr = rv
		return nil
	case "signal.target.parent.interpreter.file.name.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.parent.interpreter.file.name.length"}
	case "signal.target.parent.interpreter.file.path":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.PathnameStr"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.PathnameStr = rv
		return nil
	case "signal.target.parent.interpreter.file.path.length":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		return &eval.ErrFieldReadOnly{Field: "signal.target.parent.interpreter.file.path.length"}
	case "signal.target.parent.interpreter.file.rights":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Mode"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.Mode = uint16(rv)
		return nil
	case "signal.target.parent.interpreter.file.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.UID"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.UID = uint32(rv)
		return nil
	case "signal.target.parent.interpreter.file.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.User"}
		}
		ev.Signal.Target.Parent.LinuxBinprm.FileEvent.FileFields.User = rv
		return nil
	case "signal.target.parent.is_kworker":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.PIDContext.IsKworker"}
		}
		ev.Signal.Target.Parent.PIDContext.IsKworker = rv
		return nil
	case "signal.target.parent.is_thread":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.IsThread"}
		}
		ev.Signal.Target.Parent.IsThread = rv
		return nil
	case "signal.target.parent.pid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.PIDContext.Pid"}
		}
		ev.Signal.Target.Parent.PIDContext.Pid = uint32(rv)
		return nil
	case "signal.target.parent.ppid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.PPid"}
		}
		ev.Signal.Target.Parent.PPid = uint32(rv)
		return nil
	case "signal.target.parent.tid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.PIDContext.Tid"}
		}
		ev.Signal.Target.Parent.PIDContext.Tid = uint32(rv)
		return nil
	case "signal.target.parent.tty_name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.TTYName"}
		}
		ev.Signal.Target.Parent.TTYName = rv
		return nil
	case "signal.target.parent.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.UID"}
		}
		ev.Signal.Target.Parent.Credentials.UID = uint32(rv)
		return nil
	case "signal.target.parent.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		if ev.Signal.Target.Parent == nil {
			ev.Signal.Target.Parent = &Process{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Parent.Credentials.User"}
		}
		ev.Signal.Target.Parent.Credentials.User = rv
		return nil
	case "signal.target.pid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.PIDContext.Pid"}
		}
		ev.Signal.Target.Process.PIDContext.Pid = uint32(rv)
		return nil
	case "signal.target.ppid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.PPid"}
		}
		ev.Signal.Target.Process.PPid = uint32(rv)
		return nil
	case "signal.target.tid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.PIDContext.Tid"}
		}
		ev.Signal.Target.Process.PIDContext.Tid = uint32(rv)
		return nil
	case "signal.target.tty_name":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.TTYName"}
		}
		ev.Signal.Target.Process.TTYName = rv
		return nil
	case "signal.target.uid":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.UID"}
		}
		ev.Signal.Target.Process.Credentials.UID = uint32(rv)
		return nil
	case "signal.target.user":
		if ev.Signal.Target == nil {
			ev.Signal.Target = &ProcessContext{}
		}
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Target.Process.Credentials.User"}
		}
		ev.Signal.Target.Process.Credentials.User = rv
		return nil
	case "signal.type":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Signal.Type"}
		}
		ev.Signal.Type = uint32(rv)
		return nil
	case "splice.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.CTime"}
		}
		ev.Splice.File.FileFields.CTime = uint64(rv)
		return nil
	case "splice.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.Filesystem"}
		}
		ev.Splice.File.Filesystem = rv
		return nil
	case "splice.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.GID"}
		}
		ev.Splice.File.FileFields.GID = uint32(rv)
		return nil
	case "splice.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.Group"}
		}
		ev.Splice.File.FileFields.Group = rv
		return nil
	case "splice.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.InUpperLayer"}
		}
		ev.Splice.File.FileFields.InUpperLayer = rv
		return nil
	case "splice.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.Inode"}
		}
		ev.Splice.File.FileFields.Inode = uint64(rv)
		return nil
	case "splice.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.Mode"}
		}
		ev.Splice.File.FileFields.Mode = uint16(rv)
		return nil
	case "splice.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.MTime"}
		}
		ev.Splice.File.FileFields.MTime = uint64(rv)
		return nil
	case "splice.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.MountID"}
		}
		ev.Splice.File.FileFields.MountID = uint32(rv)
		return nil
	case "splice.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.BasenameStr"}
		}
		ev.Splice.File.BasenameStr = rv
		return nil
	case "splice.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "splice.file.name.length"}
	case "splice.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.PathnameStr"}
		}
		ev.Splice.File.PathnameStr = rv
		return nil
	case "splice.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "splice.file.path.length"}
	case "splice.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.Mode"}
		}
		ev.Splice.File.FileFields.Mode = uint16(rv)
		return nil
	case "splice.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.UID"}
		}
		ev.Splice.File.FileFields.UID = uint32(rv)
		return nil
	case "splice.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.File.FileFields.User"}
		}
		ev.Splice.File.FileFields.User = rv
		return nil
	case "splice.pipe_entry_flag":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.PipeEntryFlag"}
		}
		ev.Splice.PipeEntryFlag = uint32(rv)
		return nil
	case "splice.pipe_exit_flag":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.PipeExitFlag"}
		}
		ev.Splice.PipeExitFlag = uint32(rv)
		return nil
	case "splice.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Splice.SyscallEvent.Retval"}
		}
		ev.Splice.SyscallEvent.Retval = int64(rv)
		return nil
	case "unlink.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.CTime"}
		}
		ev.Unlink.File.FileFields.CTime = uint64(rv)
		return nil
	case "unlink.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.Filesystem"}
		}
		ev.Unlink.File.Filesystem = rv
		return nil
	case "unlink.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.GID"}
		}
		ev.Unlink.File.FileFields.GID = uint32(rv)
		return nil
	case "unlink.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.Group"}
		}
		ev.Unlink.File.FileFields.Group = rv
		return nil
	case "unlink.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.InUpperLayer"}
		}
		ev.Unlink.File.FileFields.InUpperLayer = rv
		return nil
	case "unlink.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.Inode"}
		}
		ev.Unlink.File.FileFields.Inode = uint64(rv)
		return nil
	case "unlink.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.Mode"}
		}
		ev.Unlink.File.FileFields.Mode = uint16(rv)
		return nil
	case "unlink.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.MTime"}
		}
		ev.Unlink.File.FileFields.MTime = uint64(rv)
		return nil
	case "unlink.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.MountID"}
		}
		ev.Unlink.File.FileFields.MountID = uint32(rv)
		return nil
	case "unlink.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.BasenameStr"}
		}
		ev.Unlink.File.BasenameStr = rv
		return nil
	case "unlink.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "unlink.file.name.length"}
	case "unlink.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.PathnameStr"}
		}
		ev.Unlink.File.PathnameStr = rv
		return nil
	case "unlink.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "unlink.file.path.length"}
	case "unlink.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.Mode"}
		}
		ev.Unlink.File.FileFields.Mode = uint16(rv)
		return nil
	case "unlink.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.UID"}
		}
		ev.Unlink.File.FileFields.UID = uint32(rv)
		return nil
	case "unlink.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.File.FileFields.User"}
		}
		ev.Unlink.File.FileFields.User = rv
		return nil
	case "unlink.flags":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.Flags"}
		}
		ev.Unlink.Flags = uint32(rv)
		return nil
	case "unlink.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Unlink.SyscallEvent.Retval"}
		}
		ev.Unlink.SyscallEvent.Retval = int64(rv)
		return nil
	case "utimes.file.change_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.CTime"}
		}
		ev.Utimes.File.FileFields.CTime = uint64(rv)
		return nil
	case "utimes.file.filesystem":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.Filesystem"}
		}
		ev.Utimes.File.Filesystem = rv
		return nil
	case "utimes.file.gid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.GID"}
		}
		ev.Utimes.File.FileFields.GID = uint32(rv)
		return nil
	case "utimes.file.group":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.Group"}
		}
		ev.Utimes.File.FileFields.Group = rv
		return nil
	case "utimes.file.in_upper_layer":
		rv, ok := value.(bool)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.InUpperLayer"}
		}
		ev.Utimes.File.FileFields.InUpperLayer = rv
		return nil
	case "utimes.file.inode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.Inode"}
		}
		ev.Utimes.File.FileFields.Inode = uint64(rv)
		return nil
	case "utimes.file.mode":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.Mode"}
		}
		ev.Utimes.File.FileFields.Mode = uint16(rv)
		return nil
	case "utimes.file.modification_time":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.MTime"}
		}
		ev.Utimes.File.FileFields.MTime = uint64(rv)
		return nil
	case "utimes.file.mount_id":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.MountID"}
		}
		ev.Utimes.File.FileFields.MountID = uint32(rv)
		return nil
	case "utimes.file.name":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.BasenameStr"}
		}
		ev.Utimes.File.BasenameStr = rv
		return nil
	case "utimes.file.name.length":
		return &eval.ErrFieldReadOnly{Field: "utimes.file.name.length"}
	case "utimes.file.path":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.PathnameStr"}
		}
		ev.Utimes.File.PathnameStr = rv
		return nil
	case "utimes.file.path.length":
		return &eval.ErrFieldReadOnly{Field: "utimes.file.path.length"}
	case "utimes.file.rights":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.Mode"}
		}
		ev.Utimes.File.FileFields.Mode = uint16(rv)
		return nil
	case "utimes.file.uid":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.UID"}
		}
		ev.Utimes.File.FileFields.UID = uint32(rv)
		return nil
	case "utimes.file.user":
		rv, ok := value.(string)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.File.FileFields.User"}
		}
		ev.Utimes.File.FileFields.User = rv
		return nil
	case "utimes.retval":
		rv, ok := value.(int)
		if !ok {
			return &eval.ErrValueTypeMismatch{Field: "Utimes.SyscallEvent.Retval"}
		}
		ev.Utimes.SyscallEvent.Retval = int64(rv)
		return nil
	}
	return &eval.ErrFieldNotFound{Field: field}
}
