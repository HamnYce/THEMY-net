'use client'; // This directive is necessary to ensure the component is treated as a client component

import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  useEffect(() => {
    // Redirect to the login page
    router.push("/login");
  }, [router]);

  return (
    <main>
      <p>Redirecting to login...</p>
    </main>
  );
}
