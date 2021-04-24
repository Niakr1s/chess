import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Subscription } from 'rxjs';
import { WsService } from '../ws/ws.service';

export interface Message {
  username: string;
  message: string;
}

@Injectable({
  providedIn: 'root',
})
export class ChatService {
  private eventSourceSub?: Subscription;

  private messages: Message[] = [];
  private messageSubject = new BehaviorSubject<Message[]>([]);
  public get message$(): Observable<Message[]> {
    return this.messageSubject.asObservable();
  }

  constructor(private wsService: WsService) {
    this.messageSubject.subscribe({
      next: (m) => (this.messages = m),
    });
    wsService.connected$.subscribe({
      next: (eventSource) => {
        this.clean();
        if (eventSource) {
          this.eventSourceSub = eventSource.subscribe({
            next: (e) => {
              if (e.event === 'chat:message') {
                this.messageSubject.next([
                  ...this.messages,
                  { message: e.data.message, username: e.data.username },
                ]);
              }
            },
          });
        }
      },
    });
  }

  public sendMessage(message: string): void {
    this.wsService.send({
      event: 'chat:message',
      data: { message, username: '' },
    });
  }

  private clean(): void {
    this.eventSourceSub?.unsubscribe();
    this.messageSubject.next([]);
  }
}
