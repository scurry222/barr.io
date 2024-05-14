import React from 'react'
import { Metadata } from 'next';
import Link from 'next/link';

export const metadata: Metadata = {
    title: 'About Page',
    description: 'Beginners page :)'
}

export default function ChatLayout({
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