#ifndef GO_RM_RM_H
#define GO_RM_RM_H

#include <errno.h>
#include <stdlib.h>
#include <stdint.h>
#include "./redismodule.h"

#define LOG_DEBUG   "debug"
#define LOG_VERBOSE "verbose"
#define LOG_NOTICE  "notice"
#define LOG_WARNING "warning"

int cb_cmd_func(RedisModuleCtx *ctx, RedisModuleString **argv, int argc){return cmd_func_call(ctx, argv, argc);}

int CreateCommandCall(RedisModuleCtx *ctx, const char *name,  const char *strflags, int firstkey, int lastkey, int keystep) {
    return RedisModule_CreateCommand(ctx, name, cb_cmd_func, strflags, firstkey, lastkey, keystep);
}

void CtxLog(RedisModuleCtx *ctx, int level, const char *fmt) {
    char *l;
    switch (level) {
        default:
        case 0:
            l = LOG_DEBUG;
            break;
        case 1:
            l = LOG_VERBOSE;
            break;
        case 2:
            l = LOG_NOTICE;
            break;
        case 3:
            l = LOG_WARNING;
            break;
    }
    RedisModule_Log(ctx, l, fmt);
}

int get_errno(){
    return errno;
}

#endif
