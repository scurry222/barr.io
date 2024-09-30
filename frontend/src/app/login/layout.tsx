import React from 'react'
import Link from 'next/link';

export default function LoginLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
return (
  <>
    <nav>
      <Link href="/">Back</Link>
    </nav>
    <main>
      {children}
    </main>
  </>
)
}