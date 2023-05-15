function isValidSudoku(board: string[][]): boolean {
  let seenRow = new Map<number, string[]>();
  let seenColumn = new Map<number, string[]>();
  let seenCell = new Map<string, string[]>();

  for (let r = 0; r < board.length; r++) {
    for (let c = 0; c < board.length; c++) {
      if (board[r][c] === ".") {
        continue;
      }

      const seenRowValues = seenRow.get(r);
      const seenColumnValues = seenColumn.get(c);
      const cellKey = `${Math.floor(r / 3)}-${Math.floor(c / 3)}`;
      const seenCellValues = seenCell.get(cellKey);
      const currValue = board[r][c];
      if (
        seenRowValues?.includes(currValue) ||
        seenColumnValues?.includes(currValue) ||
        seenCellValues?.includes(currValue)
      ) {
        return false;
      }

      if (seenRowValues) {
        seenRowValues.push(currValue);
      } else {
        seenRow.set(r, [currValue]);
      }

      if (seenColumnValues) {
        seenColumnValues.push(currValue);
      } else {
        seenColumn.set(c, [currValue]);
      }

      if (seenCellValues) {
        seenCellValues.push(currValue);
      } else {
        seenCell.set(cellKey, [currValue]);
      }
    }
  }

  return true;
}

const board = [
  ["5", "3", ".", ".", "7", ".", ".", ".", "."],
  ["6", ".", ".", "1", "9", "5", ".", ".", "."],
  [".", "9", "8", ".", ".", ".", ".", "6", "."],
  ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
  ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
  ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
  [".", "6", ".", ".", ".", ".", "2", "8", "."],
  [".", ".", ".", "4", "1", "9", ".", ".", "5"],
  [".", ".", ".", ".", "8", ".", ".", "7", "9"],
];

console.log("isValidSudoku", ".");
