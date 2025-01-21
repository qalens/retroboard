import { getBaseURL } from "@/lib/helper";

export async function signup(username:string,password:string){
    const resp=await fetch(`${getBaseURL()}/auth/signup`,{
        method:'POST',
        body:JSON.stringify({username,password})
    })
    if (resp.status==201){
        return (await resp.json())
    } else {
        const body=await resp.json()
        throw Error(body.message)
    }
}
export async function login(username:string,password:string){
    const resp=await fetch(`${getBaseURL()}/auth/login`,{
        method:'POST',
        body:JSON.stringify({username,password})
    })
    if (resp.status==200){
        return (await resp.json())
    } else {
        const body=await resp.json()
        throw Error(body.message)
    }
}