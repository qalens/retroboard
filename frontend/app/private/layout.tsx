import Header from "@/components/header";

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <>
            <Header loggedIn={true}/>
            {children}
        </>
    );
}
