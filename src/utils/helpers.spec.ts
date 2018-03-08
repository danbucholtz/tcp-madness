import {
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
});
