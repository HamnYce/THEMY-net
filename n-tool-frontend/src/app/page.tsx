"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import "@fortawesome/fontawesome-free/css/all.min.css";

export default function Home() {
  const router = useRouter();

  useEffect(() => {
    // Redirect to the login page at the start.
    router.push("/login");
  }, [router]);

  return (
    <main>
      <p>Redirecting to login...</p>
    </main>
  );
}
