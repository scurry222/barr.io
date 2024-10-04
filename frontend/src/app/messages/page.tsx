import TextBox from "@/components/Textbox";
import { useQuery } from "@apollo/client";
import { GET_MESSAGES } from "@/lib/graphql";

interface Message {
  id: string
  text: string
  timestamp: string
}

interface GetMessagesData {
  messages: Message[]
}

export default function Chat() {
  const { loading, error, data } = useQuery<GetMessagesData>(GET_MESSAGES);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div>
      <h1>Messages</h1>
      <ul>
        {data?.messages.map((message) => (
          <li key={message.id}>{message.text}</li>
        ))}
      </ul>
    </div>
  );
}