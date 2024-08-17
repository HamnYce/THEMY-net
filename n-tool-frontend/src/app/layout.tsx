import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "../components/header/header";
import Footer from "../components/footer/footer";
import { getOrderedFolders } from "@/utils/folderRouter";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Network Scanner Tool",
  description: "A tool for monitoring your network",
};
//TODO: Exclude the header and footer from the login page.
//The root layout for the frontend, calls the header and footer to every page.
export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const orderedFolders = getOrderedFolders();

  return (
    <html lang="en">
      <body className={`${inter.className} flex flex-col min-h-screen`}>
        <Header orderedFolders={orderedFolders} />
        <main className="flex-grow text-left">{children}</main>
        <Footer />
      </body>
    </html>
  );
}
