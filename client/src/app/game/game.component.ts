import { Component, OnInit } from '@angular/core';
import { ChessBoard } from '../services/game/chessboard';
import { GameService } from '../services/game/game.service';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css'],
})
export class GameComponent implements OnInit {
  chessboard?: ChessBoard;

  constructor(private gameService: GameService) {}

  ngOnInit(): void {
    this.gameService.chessBoard$.subscribe({
      next: (cb) => {
        this.chessboard = cb;
      },
    });
  }
}
