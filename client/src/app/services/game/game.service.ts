import { Injectable } from '@angular/core';
import { Observable, Subject, Subscription } from 'rxjs';
import { WsService } from '../ws/ws.service';
import { ChessBoard } from './chessboard';

@Injectable({
  providedIn: 'root',
})
export class GameService {
  private eventSourceSub?: Subscription;
  private chessBoard?: ChessBoard;
  private chessBoardSubject = new Subject<ChessBoard>();
  get chessBoard$(): Observable<ChessBoard> {
    return this.chessBoardSubject.asObservable();
  }

  constructor(private wsService: WsService) {
    this.wsService.connected$.subscribe({
      next: (eventSource) => {
        console.log('gameservice');
        this.eventSourceSub?.unsubscribe();
        this.eventSourceSub = eventSource?.subscribe({
          next: (e) => {
            if (e.event === 'chess:newTurn') {
              const cb = ChessBoard.decode(e.data.currentPlayer, e.data.board);
              this.chessBoard = cb;
              this.chessBoardSubject.next(cb);
            }
          },
        });
      },
    });
  }
}
