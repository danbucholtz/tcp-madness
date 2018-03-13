import {
  validateCommandFormat,
} from './commands';

import {
  INDEX,
  INDEX_NUMBER,
} from '../utils/constants';

describe('Commands', () => {
  describe('validateCommandFormat', () => {

    it('should throw when the last character is not an endline', () => {
      const expectedErrorMsg = 'Usage: Commands must end in a newline character';
      try {
        validateCommandFormat('taco|time|yo');
        fail();
      } catch (ex) {
        expect(ex.message.includes(expectedErrorMsg)).toBeTruthy();
      }

      try {
        validateCommandFormat('taco');
        fail();
      } catch (ex) {
        expect(ex.message).toEqual(expectedErrorMsg);
      }
    });

    it('should throw an error when less than 2 args or more than 3 args', () => {

      const expectedErrorMsg = 'Usage: command|package|dep1,dep2,dep3,...';
      try {
        validateCommandFormat('\n');
        fail();
      } catch (ex) {
        expect(ex.message).toEqual(expectedErrorMsg);
      }

      try {
        validateCommandFormat('taco\n');
        fail();
      } catch (ex) {
        expect(ex.message).toEqual(expectedErrorMsg);
      }

      try {
        validateCommandFormat('taco|\n');
        fail();
      } catch (ex) {
        expect(ex.message).toEqual(expectedErrorMsg);
      }

      try {
        validateCommandFormat('taco|time\n');
        fail();
      } catch (ex) {
        expect(ex.message).toEqual(expectedErrorMsg);
      }

      try {
        validateCommandFormat('taco|time|for|everyone\n');
        fail();
      } catch (ex) {
        expect(ex.message).toEqual(expectedErrorMsg);
      }
    });

    it('should throw when the first argument is invalid', () => {
      const expectedErrorMsg = 'First argument must be';
      try {
        validateCommandFormat('taco|time|\n');
        fail();
      } catch (ex) {
        expect(ex.message.includes(expectedErrorMsg)).toBeTruthy();
      }
    });

    it('should return an instance of command interface', () => {
      const result = validateCommandFormat(`${INDEX}|ls|\n`);
      expect(result.dependencies).toBeTruthy();
      expect(result.dependencies.length).toBe(0);
      expect(result.type).toBe(INDEX_NUMBER);
      expect(result.packageName).toBe('ls');
    });

    it('should return an instance of command interface with dependencies', () => {
      const result = validateCommandFormat(`${INDEX}|ls|cd,cat,man,tsc,node,npm\n`);
      expect(result.dependencies.length).toBe(6);
      expect(result.dependencies[0]).toBe('cd');
      expect(result.dependencies[1]).toBe('cat');
      expect(result.dependencies[2]).toBe('man');
      expect(result.dependencies[3]).toBe('tsc');
      expect(result.dependencies[4]).toBe('node');
      expect(result.dependencies[5]).toBe('npm');
      expect(result.type).toBe(INDEX_NUMBER);
      expect(result.packageName).toBe('ls');
    });

    it('should account for weird spacing, etc in dependencies', () => {
      const result = validateCommandFormat(`${INDEX}|ls|cd, cat, man,tsc, node,npm\n`);
      expect(result.dependencies.length).toBe(6);
      expect(result.dependencies[0]).toBe('cd');
      expect(result.dependencies[1]).toBe('cat');
      expect(result.dependencies[2]).toBe('man');
      expect(result.dependencies[3]).toBe('tsc');
      expect(result.dependencies[4]).toBe('node');
      expect(result.dependencies[5]).toBe('npm');
    });

    it('should account for empty csv entry in the dependencies', () => {
      const knownErrorMsg = 'Usage: One or more invalid dependency entries';
      try {
        validateCommandFormat(`${INDEX}|ls|cd,,man, tsc,,node, npm\n`);
        fail();
      } catch (ex) {
        expect(ex.message.includes(knownErrorMsg)).toBeTruthy();
      }
    });

    it('should throw an error when command has a space in it', () => {
      const knownErrorMsg = 'Usage: Commands must be a valid unix command format';
      try {
        validateCommandFormat(`INDEX|emacs elisp|\n`);
        fail();
      } catch (ex) {
        expect(ex.message.includes(knownErrorMsg)).toBeTruthy();
      }
    });
  });
});

