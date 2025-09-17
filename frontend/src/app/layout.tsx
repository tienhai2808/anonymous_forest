import type { Metadata } from "next";
import { IBM_Plex_Mono } from "next/font/google";
import "./globals.css";
import SideBar from "@/components/SideBar";

const ibmPlexMono = IBM_Plex_Mono({
  subsets: ["vietnamese"],
  variable: "--font-ibm",
  display: "swap",
  weight: ["100", "300", "400", "500", "700"],
  preload: true,
});

export const metadata: Metadata = {
  title: "Chốn an yên",
  description: "Nơi trút bầu tâm sự",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={ibmPlexMono.className}>
        <div className="flex min-h-screen">
          <SideBar/>
          <div className="flex flex-col flex-1 bg-white">
            {children}
          </div>
        </div>
      </body>
    </html>
  );
}
