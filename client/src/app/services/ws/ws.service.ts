import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Subject, Subscription } from 'rxjs';
import { webSocket } from 'rxjs/webSocket';
import { AuthUsernameEvent, Event } from './events';

type EventSource = Subject<Event> | null;

@Injectable({
  providedIn: 'root',
})
export class WsService {
  private connectedSubject = new BehaviorSubject<EventSource>(null);
  private lastSubject: EventSource = null;

  public get connected$(): Observable<EventSource> {
    return this.connectedSubject.asObservable();
  }

  private wsSub?: Subscription;

  constructor() {
    this.connected$.subscribe({
      next: (s) => (this.lastSubject = s),
    });
  }

  connect(username: string): Observable<Event> {
    this.clean();
    const wsUrl =
      (window.location.protocol === 'https:' ? 'wss://' : 'ws://') +
      window.location.host +
      '/ws';
    console.log('connecting to ', wsUrl);
    const ws = webSocket<Event>({ url: wsUrl });
    this.wsSub = ws.subscribe({
      next: (e) => console.log('got event', e),
      complete: () => this.clean(),
      error: () => this.clean(),
    });

    const authUsernameEvent: AuthUsernameEvent = {
      event: 'auth:username',
      data: { username },
    };
    ws.next(authUsernameEvent);

    this.connectedSubject.next(ws);
    return ws;
  }

  send(event: Event): void {
    console.log('sending', event, this.lastSubject);
    this.lastSubject?.next(event);
  }

  private clean(): void {
    this.lastSubject?.complete();
    this.connectedSubject.next(null);
    this.wsSub?.unsubscribe();
  }
}
