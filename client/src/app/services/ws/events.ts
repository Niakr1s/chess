export type Event =
  | ChatMessageEvent
  | ChatUserJoinEvent
  | AuthUsernameEvent
  | ChessNewTurnEvent;

export interface ChatMessageEvent {
  event: 'chat:message';
  data: {
    message: string;
    username: string;
  };
}

export interface ChatUserJoinEvent {
  event: 'chat:userJoin';
  data: {
    username: string;
  };
}

export interface AuthUsernameEvent {
  event: 'auth:username';
  data: {
    username: string;
  };
}

export interface ChessNewTurnEvent {
  event: 'chess:newTurn';
  data: {
    board: string;
    currentPlayer: number;
  };
}
