import { decode } from 'base64-arraybuffer';
import { Color, Figure } from './figure';

export class ChessBoard {
  currentPlayer: Color = 'white';
  board!: (Figure | null)[][];

  static decode(currentPlayer: number, encodedBoard: string): ChessBoard {
    const res = new ChessBoard();
    res.currentPlayer = currentPlayer === 0 ? 'white' : 'black';

    const bytes = decode(encodedBoard);
    const view = new Uint8Array(bytes);
    console.log('bytes', bytes);

    const board: (Figure | null)[][] = [];

    for (let row = 0; row < 8; row++) {
      const rowArr: (Figure | null)[] = [];
      for (let col = 0; col < 8; col++) {
        const index = row * 8 + col;
        rowArr.push(Figure.decode(view[index]));
      }
      board.push(rowArr);
    }

    res.board = board;
    return res;
  }
}
