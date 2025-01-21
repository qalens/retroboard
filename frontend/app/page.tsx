import { redirect } from "next/navigation";
import { cookies } from "next/headers";
export default async function Home() {
  const store = await cookies()
  const session = store.get('token')
  if (!session){
    redirect("/login")
  } else {
    redirect("/home")
  }
}
