

export function debug(content: string, ...other: any[]) {
  if (shouldLog(0)) {
    return logImpl(content, ...other);
  }
}

export function minimal(content: string, ...other: any[]) {
  if (shouldLog(1)) {
    return logImpl(content, ...other);
  }
}

function logImpl(content: string, ...other: any[]) {
  // in a real app, it would make way more sense to write to a stream
  console.log(`[${(new Date()).toString()}] ${content}: `, other);
}

export function shouldLog(requestedLogLevel: number): boolean {
  const applicationLogLevel = getLogLevel(process.env.LOG_LEVEL);
  return requestedLogLevel > applicationLogLevel;
}

function getLogLevel(input: string) {
  if (input === DEBUG) {
    return 0;
  } else if (input === MINIMAL) {
    return 1;
  }
  return 0;
}

export const DEBUG = "DEBUG";
export const MINIMAL = "MINIMAL";