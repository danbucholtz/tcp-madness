import { processCommand } from './process-command';

import { wipeDataStore } from './data-access/data-access';

import {
  INDEX_NUMBER,
  REMOVE_NUMBER,
  REMOVE
} from './utils/constants';

describe('process commands', () => {
  describe('processCommand', () => {

    beforeEach(() => {
      wipeDataStore();
    });

    it('should return false if it cannot process an index command', async () => {
      const result = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'ls',
        dependencies: ['someCommand', 'someOtherCommand']
      });
      expect(result).toBe(false);
    });

    it('should return true when it can process an index command', async () => {
      const result = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'ls',
        dependencies: []
      });
      expect(result).toBe(true);
    });

    it('should handle multiple correct index commands', async () => {
      const firstResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'ls',
        dependencies: []
      });

      expect(firstResult).toBe(true);

      const secondResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'cd',
        dependencies: ['ls']
      });

      expect(secondResult).toBe(true);

      const thirdResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'pwd',
        dependencies: ['ls', 'cd']
      });

      expect(thirdResult).toBe(true);
    });

    it('should handle an incorrect command correctly among set of correct index commands', async () => {
      const firstResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'ls',
        dependencies: []
      });

      expect(firstResult).toBe(true);

      const secondResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'cd',
        dependencies: ['ls']
      });

      expect(secondResult).toBe(true);

      const thirdResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'pwd',
        dependencies: ['ls', 'cd']
      });

      expect(thirdResult).toBe(true);

      const fourthResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'blah',
        dependencies: ['doesntExistYet']
      });

      expect(fourthResult).toBe(false);

      const fifthResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'cat',
        dependencies: ['cd', 'tacobell']
      });

      expect(fifthResult).toBe(false);

      const sixthResult = await processCommand({
        type: INDEX_NUMBER,
        packageName: 'node',
        dependencies: []
      });

      expect(sixthResult).toBe(true);
    });

    it('should return true if it can process a remove command', async () => {
      // arrange default data
      await processCommand({
        type: INDEX_NUMBER,
        packageName: 'ls',
        dependencies: ['someCommand', 'someOtherCommand']
      });

      // act
      const result = await processCommand({
        type: REMOVE_NUMBER,
        packageName: 'ls',
        dependencies: []
      });

      // assert
      expect(result).toBe(true);
    });

    it('should return true if there isn\'t a command to remove', async () => {
      const result = await processCommand({
        type: REMOVE_NUMBER,
        packageName: 'ls',
        dependencies: []
      });

      // assert
      expect(result).toBe(true);
    });

    it('should return false if it cannot execute a remove command', async () => {
      // arrange default data
      
      await processCommand({
        type: INDEX_NUMBER,
        packageName: 'ls',
        dependencies: []
      });

      await processCommand({
        type: INDEX_NUMBER,
        packageName: 'someCommand',
        dependencies: ['ls']
      });

      // act
      const result = await processCommand({
        type: REMOVE_NUMBER,
        packageName: 'ls',
        dependencies: []
      });

      // assert
      expect(result).toBe(false);
    });
  });

});
