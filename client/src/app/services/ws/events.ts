export type Event = ChatMessageEvent | ChatUserJoinEvent | AuthUsernameEvent;

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
