

export function reduceBools(list: boolean[]): boolean {
  const blah = list.every((value: boolean) => !!value);
  return blah;
}