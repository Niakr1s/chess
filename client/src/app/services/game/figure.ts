export const FigureKing = 'king';
export const FigureQueen = 'queen';
export const FigureBishop = 'bishop';
export const FigureKnight = 'knight';
export const FigureRook = 'rook';
export const FigurePawn = 'pawn';

export type FigureType =
  | 'king'
  | 'queen'
  | 'bishop'
  | 'knight'
  | 'rook'
  | 'pawn';

export type Color = 'white' | 'black';

export class Figure {
  type: FigureType;
  color: Color;
  wasMoved: boolean;

  constructor(type: FigureType, color: Color, wasMoved: boolean = false) {
    this.type = type;
    this.color = color;
    this.wasMoved = wasMoved;
  }

  static decode(b: number): Figure | null {
    if (b === 0) {
      return null;
    }
    let type: FigureType;
    // tslint:disable-next-line: no-bitwise
    if (b & 0b00000001) {
      type = 'king';
      // tslint:disable-next-line: no-bitwise
    } else if (b & 0b00000010) {
      type = 'queen';
      // tslint:disable-next-line: no-bitwise
    } else if (b & 0b00000100) {
      type = 'bishop';
      // tslint:disable-next-line: no-bitwise
    } else if (b & 0b00001000) {
      type = 'knight';
      // tslint:disable-next-line: no-bitwise
    } else if (b & 0b00010000) {
      type = 'rook';
    } else {
      type = 'pawn';
    }
    // tslint:disable-next-line: no-bitwise
    const color: Color = b & 0b01000000 ? 'black' : 'white';
    // tslint:disable-next-line: no-bitwise
    const wasMoved: boolean = !!(b & 0b10000000);
    return new Figure(type, color, wasMoved);
  }

  toString(): string {
    if (this.color === 'white') {
      switch (this.type) {
        case 'king':
          return '♔';
        case 'queen':
          return '♕';
        case 'bishop':
          return '♗';
        case 'knight':
          return '♘';
        case 'rook':
          return '♖';
        case 'pawn':
          return '♙';
      }
    } else {
      switch (this.type) {
        case 'king':
          return '♚';
        case 'queen':
          return '♛';
        case 'bishop':
          return '♝';
        case 'knight':
          return '♞';
        case 'rook':
          return '♜';
        case 'pawn':
          return '♟';
      }
    }
  }
}
