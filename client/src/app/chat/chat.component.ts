import { Component, OnInit } from '@angular/core';
import { ChatService, Message } from '../services/chat/chat.service';
import { WsService } from '../services/ws/ws.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css'],
})
export class ChatComponent implements OnInit {
  messages: Message[] = [];
  username = '';
  message = '';
  connected = false;

  constructor(private wsService: WsService, private chatService: ChatService) {}

  ngOnInit(): void {
    this.chatService.message$.subscribe({
      next: (m) => {
        this.messages = m;
      },
    });
    this.wsService.connected$.subscribe({
      next: (e) => (this.connected = !!e),
    });
  }

  onConnect(): void {
    this.wsService.connect(this.username);
  }

  onMessage(): void {
    this.chatService.sendMessage(this.message);
  }
}
