import { redirect } from 'next/navigation';
import { cookies } from "next/headers";
export default async function Contact({ searchParams }: { searchParams: Promise<{ q?: string }> }) {
  const store = await cookies()
  const token = store.get('token')
  if (!token) {
    redirect("/login")
  } else {
    return (<>
      <div className="container mx-auto max-w-7xl p-6 flex-grow border my-2 rounded-xl ">
        Welcome
      </div>
    </>
    );
  }
}
