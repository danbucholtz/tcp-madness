

export function reduceBools(list: boolean[]): boolean {
  const blah = list.every((value: boolean) => !!value);
  return blah;
}

export function containsOnlyLetters(input: string): boolean {
  const regex = /^[A-Za-z]+$/;
  const result = regex.exec(input);
  return !!result;
}