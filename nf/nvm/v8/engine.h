// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see
// <http://www.gnu.org/licenses/>.
//

#ifndef _NEBULAS_NF_NVM_V8_ENGINE_H_
#define _NEBULAS_NF_NVM_V8_ENGINE_H_
// #include <v8.h>

#if BUILDING_DLL
#define EXPORT __attribute__((__visibility__("default")))
#else
#define EXPORT
#endif

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

#include <stddef.h>
#include <stdint.h>
#include <string.h>
#include <stdbool.h>

enum LogLevel {
  DEBUG = 1,
  WARN = 2,
  INFO = 3,
  ERROR = 4,
};

enum OptType {
  INSTRUCTION     = 1,
  INSTRUCTIONTS  = 2,
  RUNSCRIPT       = 3,
};

// log
typedef void (*LogFunc)(int level, const char *msg);
EXPORT const char *GetLogLevelText(int level);
EXPORT void InitializeLogger(LogFunc f);

// event.
typedef void (*EventTriggerFunc)(void *handler, const char *topic,
                                 const char *data, size_t *counterVal);
EXPORT void InitializeEvent(EventTriggerFunc trigger);

// storage
typedef char *(*StorageGetFunc)(void *handler, const char *key,
                                size_t *counterVal);
typedef int (*StoragePutFunc)(void *handler, const char *key, const char *value,
                              size_t *counterVal);
typedef int (*StorageDelFunc)(void *handler, const char *key,
                              size_t *counterVal);
EXPORT void InitializeStorage(StorageGetFunc get, StoragePutFunc put,
                              StorageDelFunc del);

// blockchain
typedef char *(*GetTxByHashFunc)(void *handler, const char *hash,
                                 size_t *counterVal);
typedef char *(*GetAccountStateFunc)(void *handler, const char *address,
                                     size_t *counterVal);
typedef int (*TransferFunc)(void *handler, const char *to, const char *value,
                            size_t *counterVal);
typedef int (*VerifyAddressFunc)(void *handler, const char *address,
                                 size_t *counterVal);

EXPORT void InitializeBlockchain(GetTxByHashFunc getTx,
                                 GetAccountStateFunc getAccount,
                                 TransferFunc transfer,
                                 VerifyAddressFunc verifyAddress);

// version
EXPORT char *GetV8Version();

// require callback.
typedef char *(*RequireDelegate)(void *handler, const char *filename,
                                 size_t *lineOffset);
EXPORT void InitializeRequireDelegate(RequireDelegate delegate);

typedef struct V8EngineStats {
  size_t count_of_executed_instructions;
  size_t total_memory_size;
  size_t total_heap_size;
  size_t total_heap_size_executable;
  size_t total_physical_size;
  size_t total_available_size;
  size_t used_heap_size;
  size_t heap_size_limit;
  size_t malloced_memory;
  size_t peak_malloced_memory;
  size_t total_array_buffer_size;
  size_t peak_array_buffer_size;
} V8EngineStats;

typedef struct V8Engine {
  void *isolate;
  void *allocator;
  size_t limits_of_executed_instructions;
  size_t limits_of_total_memory_size;
  int is_requested_terminate_execution;
  int testing;
  uintptr_t lcs;  
  uintptr_t gcs;
  enum OptType opt;  
  int lineOffset;
  int ret;  //output
  int allowUsage;
  V8EngineStats stats;
  const char *source;
  char *result; //output
  bool isRunEnd;
  
  // V8Engine() {
  //   // memset(this, 0x00, sizeof(V8Engine));
  // }
  // void *context;
} V8Engine;

EXPORT void Initialize();
EXPORT void Dispose();

EXPORT V8Engine *CreateEngine();

EXPORT int RunScriptSource(char **result, V8Engine *e, const char *source,
                           int source_line_offset, uintptr_t lcsHandler,
                           uintptr_t gcsHandler);

EXPORT char *InjectTracingInstructions(V8Engine *e, const char *source,
                                       int *source_line_offset,
                                       int strictDisallowUsage);

EXPORT char *TranspileTypeScriptModule(V8Engine *e, const char *source,
                                       int *source_line_offset);

EXPORT int IsEngineLimitsExceeded(V8Engine *e);

EXPORT void ReadMemoryStatistics(V8Engine *e);

EXPORT void TerminateExecution(V8Engine *e);

EXPORT void DeleteEngine(V8Engine *e);

EXPORT void ExecuteLoop(const char *file);

// EXPORT V8Engine *CreateThreadEngine();
// EXPORT void DeleteThreadEngine(V8Engine *e);

// EXPORT void ExecuteLoopInIsolate(V8Engine *e);
EXPORT char *InjectTracingInstructionsThread(V8Engine *e, const char *source,
                                int *source_line_offset,
                                int strictDisallowUsage);
EXPORT char *TranspileTypeScriptModuleThread(V8Engine *e, const char *source,
                                int *source_line_offset);
int RunScriptSourceThread(char **result, V8Engine *e, const char *source,
                    int source_line_offset, uintptr_t lcsHandler,
                    uintptr_t gcsHandler);
EXPORT void RunScriptThread(V8Engine *e);
EXPORT void SetRunScriptArgs(V8Engine *e, int opt, const char *source, int lineOffset, int allowUsage);
EXPORT void DecoratorOutPut(V8Engine *e);

#ifdef __cplusplus
}
#endif // __cplusplus

#endif // _NEBULAS_NF_NVM_V8_ENGINE_H_
