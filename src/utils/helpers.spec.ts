import {
  containsOnlyLetters,
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

  describe('containsOnlyLetters', () => {
    it('should return false if a non letter is in the string', () => {
      const resultOne = containsOnlyLetters('abc ');
      expect(resultOne).toBe(false);

      const resultTwo = containsOnlyLetters('abc123');
      expect(resultTwo).toBe(false);

      const resultThree = containsOnlyLetters('abc?');
      expect(resultThree).toBe(false);

      const resultFour = containsOnlyLetters('');
      expect(resultFour).toBe(false);

      const resultFive = containsOnlyLetters(' ');
      expect(resultFive).toBe(false);

      const resultSix = containsOnlyLetters('(');
      expect(resultSix).toBe(false);

      const resultSeven = containsOnlyLetters('ABC ');
      expect(resultSeven).toBe(false);
    });

    it('should return true if only letters are in the string', () => {
      const resultOne = containsOnlyLetters('abc');
      expect(resultOne).toBe(true);

      const resultTwo = containsOnlyLetters('q');
      expect(resultTwo).toBe(true);

      const resultThree = containsOnlyLetters('ABC');
      expect(resultThree).toBe(true);

      const resultFour = containsOnlyLetters('Q');
      expect(resultFour).toBe(true);
    });
  });
});
