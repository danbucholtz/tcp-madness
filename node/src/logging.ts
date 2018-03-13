const DEBUG_LEVEL = 'debug'
const DEBUG_LEVEL_NUMBER = 1
const WARN_LEVEL = 'warn'
const WARN_LEVEL_NUMBER = 2

let loggingNumber = WARN_LEVEL_NUMBER

const logLevelMap = new Map<string, number>()
logLevelMap.set(DEBUG_LEVEL, DEBUG_LEVEL_NUMBER);
logLevelMap.set(WARN_LEVEL, WARN_LEVEL_NUMBER);

export function initLogger() {
  let levelString = process.env['LOGGING_LEVEL'];
  console.log('Logger Environment variable String is set to: ', levelString);
  if (levelString !== DEBUG_LEVEL) {
    levelString = WARN_LEVEL;
  }

  loggingNumber = levelString === DEBUG_LEVEL ? DEBUG_LEVEL_NUMBER : WARN_LEVEL_NUMBER;

  console.log('Logger level set to: ', levelString);
}

export function debug(input: string) {
  if (loggingNumber <= DEBUG_LEVEL_NUMBER) {
    return loggerImpl(input);
  }
}

export function warn(input: string) {
  if (loggingNumber <= WARN_LEVEL_NUMBER) {
    return loggerImpl(input);
  }
}

function loggerImpl(input: string) {
  return console.log(input);
}