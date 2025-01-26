import { Button } from "@heroui/button";
import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from "@heroui/navbar";
import Link from "next/link";

export default function Header({loggedIn}:{loggedIn:boolean}){
    return <Navbar>
        <NavbarBrand>
            <p>Retroboard</p>
        </NavbarBrand>
        <NavbarContent justify="end">
        {!loggedIn && <><NavbarItem className="hidden lg:flex">
          <Link href="/public/login">Login</Link>
        </NavbarItem>
        <NavbarItem>
          <Button as={Link} color="primary" href="/public/signup" variant="flat">
            Sign Up
          </Button>
        </NavbarItem>
        </>}
        {loggedIn && <NavbarItem>
          <Button as={Link} color="primary" href="/api/auth/logout" variant="flat">
            Logout
          </Button>
        </NavbarItem>
        }
      </NavbarContent>
    </Navbar>
}