import { Component, OnInit } from '@angular/core';
import { timer } from 'rxjs';
import { webSocket } from 'rxjs/webSocket';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  ngOnInit(): void {
    const wsUrl =
      (window.location.protocol === 'https:' ? 'wss://' : 'ws://') +
      window.location.host +
      '/ws';
    console.log('connecting to ', wsUrl);
    const ws = webSocket({ url: wsUrl });
    ws.subscribe({
      next: (e) => {
        console.log(e);
      },
    });
    ws.next({
      event: 'auth:username',
      data: { username: 'username' },
    });
    timer(200, 1000).subscribe({
      next: () => {
        ws.next({
          event: 'chat:message',
          data: { message: 'message' },
        });
      },
    });
  }
}
