'use client'
import { signup } from "@/services/user";
import { Button } from "@heroui/button";
import { Card, CardBody, CardFooter, CardHeader } from "@heroui/card";
import { Input } from "@heroui/input";
import Link from "next/link";
import { useState } from "react";
import {redirect} from 'next/navigation'
import { useToast } from "@/hooks/use-toast";
export default function SignUp() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const {toast}=useToast()
    const [confirmPassword, setConfirmPassword] = useState('')
    return <Card className="w-[400px] mx-auto">
        <CardHeader>Sign Up</CardHeader>
        <CardBody className="flex flex-col gap-2">
            <Input placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
            <Input placeholder="Password" type="password" value={password} onChange={e => setPassword(e.target.value)} />
            <Input placeholder="Confirm Password" type="password" value={confirmPassword} onChange={e => setConfirmPassword(e.target.value)} />
        </CardBody>
        <CardFooter className="flex flex-row items-center justify-right gap-2">
            <Link href="/login"><Button>Login</Button></Link>
            <Button className="grow" color="primary" onPress={() => {
                if (password == confirmPassword) {
                    signup(username, password).then(token => {
                        const expiryDate = new Date();
                        expiryDate.setDate(expiryDate.getDate() + 7); // Expires in 7 days
                        document.cookie = `token=${token}; path=/; secure; samesite=strict; expires=${expiryDate.toUTCString()}`;
                        return true
                    }).catch(e => {
                        toast({
                            title:'Failure',
                            description: e.message,
                            variant:'destructive'
                        })
                        return false
                    }).then((success)=>{
                        if (success)
                            redirect('/home')
                    })
                }
            }}>Sign Up</Button>
        </CardFooter>
    </Card>
}