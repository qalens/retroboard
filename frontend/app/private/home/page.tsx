import { redirect } from 'next/navigation';
import { cookies } from "next/headers";
import Link from 'next/link';
export default async function Contact({ searchParams }: { searchParams: Promise<{ q?: string }> }) {
  const store = await cookies()
  const token = store.get('token')
  if (!token) {
    redirect("/public/login")
  } else {
    return (<>
      <div className="container mx-auto max-w-7xl p-6 flex-grow border my-2 rounded-xl ">
        Welcome
      </div>
    </>
    );
  }
}
