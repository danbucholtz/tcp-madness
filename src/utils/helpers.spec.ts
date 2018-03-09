import {
  isValidUnixCommand,
  reduceBools
} from './helpers';

describe('helpers', () => {
  describe('reduceBools', () => {
    it('should return false when there is a false in the list', () => {
      const result = reduceBools([true, false, true]);
      expect(result).toBeFalsy();
    });

    it('should return true when all values are true', () => {
      const result = reduceBools([true, true, true]);
      expect(result).toBeTruthy();
    });
  });

  describe('isValidUnixCommand', () => {
    it('should return false if a non alphanumeric, underscore or hyphen is in the string', () => {
      const resultOne = isValidUnixCommand('abc ');
      expect(resultOne).toBe(false);

      const resultThree = isValidUnixCommand('abc?');
      expect(resultThree).toBe(false);

      const resultFour = isValidUnixCommand('');
      expect(resultFour).toBe(false);

      const resultFive = isValidUnixCommand(' ');
      expect(resultFive).toBe(false);

      const resultSix = isValidUnixCommand('(');
      expect(resultSix).toBe(false);

      const resultSeven = isValidUnixCommand('ABC ');
      expect(resultSeven).toBe(false);
    });

    it('should return true if string is only alphanumeric, underscore or hyphen', () => {
      const resultOne = isValidUnixCommand('abc');
      expect(resultOne).toBe(true);

      const resultTwo = isValidUnixCommand('q');
      expect(resultTwo).toBe(true);

      const resultThree = isValidUnixCommand('ABC');
      expect(resultThree).toBe(true);

      const resultFour = isValidUnixCommand('Q');
      expect(resultFour).toBe(true);

      const resultFive = isValidUnixCommand('abc123');
      expect(resultFive).toBe(true);

      const resultSix = isValidUnixCommand('abc_123');
      expect(resultSix).toBe(true);

      const resultSeven = isValidUnixCommand('abc-123');
      expect(resultSeven).toBe(true);

      const resultEight = isValidUnixCommand('abc+123');
      expect(resultEight).toBe(true);
    });
  });
});
