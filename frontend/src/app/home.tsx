import NavigateToPage from "./components/NavigateToPage";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col justify-between p-24">
      <header className="flex justify-between">
        <h1>WELCOME</h1>
        <NavigateToPage pageName="login" title="Log In" />
      </header>
      <nav>
        <NavigateToPage pageName="messages" title="Messages" />
      </nav>
    </main>
  );
}
