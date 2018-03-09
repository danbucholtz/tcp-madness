

export function reduceBools(list: boolean[]): boolean {
  const blah = list.every((value: boolean) => !!value);
  return blah;
}

export function isValidUnixCommand(input: string): boolean {
  const regex = /^[a-zA-Z0-9_+-]+$/;
  const result = regex.exec(input);
  return !!result;
}