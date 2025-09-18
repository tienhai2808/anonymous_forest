import type { Metadata } from "next";
import { IBM_Plex_Mono } from "next/font/google";
import "./globals.css";
import SideBar from "@/components/SideBar";
import ThemeProvider from "@/components/theme-provider";

const ibmPlexMono = IBM_Plex_Mono({
  subsets: ["vietnamese"],
  variable: "--font-ibm",
  display: "swap",
  weight: ["100", "300", "400", "500", "700"],
  preload: true,
});

export const metadata: Metadata = {
  title: "Chốn An Yên",
  description: "Nơi trút bầu tâm sự",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={ibmPlexMono.className}>
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem disableTransitionOnChange>
          <div className="flex min-h-screen">
            <SideBar />
            <div className="flex flex-col flex-1">{children}</div>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
