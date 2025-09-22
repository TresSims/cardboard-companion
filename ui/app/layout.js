import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Header from '../components/header'
import Footer from '../components/footer'

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata = {
  title: "Cardboard Companion",
  description: "A discrod bot, lovingly gopher flavored",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased flex flex-col w-full min-h-[100vh]`}
      >
        <Header />
        <div className="flex-grow flex flex-row g-[10px] p-[10px]">
          {children}
        </div>
        <Footer />
      </body>
    </html>
  );
}
