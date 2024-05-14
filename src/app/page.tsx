import Link from "next/link";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <header className="flex flex-col items-center">
        <h1>HII WELCOME</h1>
        <img src="hii.webp"alt="HII" />
      </header>
      <nav>
        <Link href='/chat'>TO THE CHAT</Link>
      </nav>
    </main>
  );
}
