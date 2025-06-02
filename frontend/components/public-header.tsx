"use client"
import { useState, useEffect } from "react"
import type React from "react"

import Link from "next/link"
import { useRouter, usePathname } from "next/navigation"
import { Button } from "@/components/ui/button"
import { Sheet, SheetContent, SheetTrigger, SheetHeader, SheetTitle } from "@/components/ui/sheet"
import {
    Menu,
    Settings,
    User,
    LogOut,
} from "lucide-react"
import { cn } from "@/lib/utils"
import { useAuth } from "@/context/auth-context"
import { useAuthActions } from "@/hooks/use-auth-actions"
import GradeFlowLogo from "./logo"
import { navigationItems } from "@/lib/constants/navigation"

type PageRoute = '/landing-page' | '/login' | '/register'


export default function Header({ currentPage }: { currentPage: PageRoute }) {
    const { logout } = useAuthActions()
    const { user, isAuthenticated } = useAuth()
    const router = useRouter()
    const pathname = usePathname()
    const [isScrolled, setIsScrolled] = useState(false)

    useEffect(() => {
        const handleScroll = () => {
            setIsScrolled(window.scrollY > 10)
        }
        window.addEventListener("scroll", handleScroll)
        return () => window.removeEventListener("scroll", handleScroll)
    }, [])

    const handleLogout = () => {
        logout()
        router.push("/")
    }

    return (
        <header
            className={cn(
                "sticky top-0 z-50 w-full border-b transition-all duration-200",
                isScrolled ? "bg-background/80 backdrop-blur-md shadow-sm" : "bg-background/95 backdrop-blur-sm",
                "dark:bg-background/90",
            )}
        >
            < div className="max-w-7xl mx-auto flex h-16 items-center justify-between px-4">
                <Link
                    href={isAuthenticated ? "/dashboard" : "/"}
                    className="flex items-center space-x-2 transition-opacity hover:opacity-80"
                >
                    <GradeFlowLogo width={200} height={100} />
                </Link>


                {currentPage === '/landing-page' && (
                    <nav className="hidden md:flex items-center space-x-6">
                        <Link
                            href="#recursos"
                            className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
                        >
                            Recursos
                        </Link>
                        <Link
                            href="#precos"
                            className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
                        >
                            Preços
                        </Link>
                        <Link
                            href="#sobre"
                            className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
                        >
                            Sobre
                        </Link>
                        <Link
                            href="#contato"
                            className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
                        >
                            Contato
                        </Link>
                    </nav>
                )}


                {!isAuthenticated ?
                    <div className="hidden md:flex items-center space-x-2">
                        <Button variant="ghost" asChild>
                            <Link href="/login">Log in</Link>
                        </Button>
                        <Button asChild>
                            <Link href="/register">Sign up</Link>
                        </Button>
                    </div> : <div>
                        <Button onClick={handleLogout} variant='outline'>
                            Sair
                            <LogOut />
                        </Button>
                    </div>}

                {/* Mobile Menu */}
                <Sheet>
                    <SheetTrigger asChild>
                        <Button variant="ghost" size="icon" className="md:hidden h-9 w-9">
                            <Menu className="h-4 w-4" />
                            <span className="sr-only">Open menu</span>
                        </Button>
                    </SheetTrigger>
                    <SheetContent side="right" className="w-[300px] sm:w-[400px]">
                        <SheetHeader>
                            <SheetTitle>Menu</SheetTitle>
                        </SheetHeader>
                        <div className="mt-6 flex flex-col space-y-4">
                            {isAuthenticated ? (
                                <>
                                    {/* User Info */}
                                    <div className="flex items-center space-x-3 rounded-lg border p-3">
                                        <div className="flex-1 min-w-0">
                                            <p className="text-sm font-medium truncate">{user?.username}</p>
                                            <p className="text-xs text-muted-foreground truncate">{user?.email}</p>
                                        </div>
                                    </div>


                                    {/* Mobile Navigation */}
                                    <nav className="flex flex-col space-y-1">
                                        {navigationItems.map((item) => {
                                            const isActive = pathname === item.href
                                            return (
                                                <Link
                                                    key={item.href}
                                                    href={item.href}
                                                    className={cn(
                                                        "flex items-center space-x-3 rounded-md px-3 py-2 text-sm font-medium transition-colors",
                                                        isActive
                                                            ? "bg-accent text-accent-foreground"
                                                            : "text-muted-foreground hover:bg-accent hover:text-accent-foreground",
                                                    )}
                                                >
                                                    <item.icon className="h-4 w-4" />
                                                    <span>{item.label}</span>
                                                </Link>
                                            )
                                        })}
                                    </nav>

                                    {/* Mobile Actions */}
                                    <div className="space-y-2 border-t pt-4">
                                        <Button variant="outline" className="w-full justify-start" asChild>
                                            <Link href="/profile">
                                                <User className="mr-2 h-4 w-4" />
                                                Profile
                                            </Link>
                                        </Button>
                                        <Button variant="outline" className="w-full justify-start" asChild>
                                            <Link href="/settings">
                                                <Settings className="mr-2 h-4 w-4" />
                                                Settings
                                            </Link>
                                        </Button>
                                        <Button variant="outline" className="w-full justify-start text-red-600" onClick={handleLogout}>
                                            <LogOut className="mr-2 h-4 w-4" />
                                            Log out
                                        </Button>
                                    </div>
                                </>
                            ) : (
                                <>
                                    {/* Public Mobile Navigation */}
                                    <nav className="flex flex-col space-y-2">
                                        <Link
                                            href="#recursos"
                                            className="text-lg font-medium text-muted-foreground hover:text-foreground transition-colors"
                                        >
                                            Recursos
                                        </Link>
                                        <Link
                                            href="#precos"
                                            className="text-lg font-medium text-muted-foreground hover:text-foreground transition-colors"
                                        >
                                            Preços
                                        </Link>
                                        <Link
                                            href="#sobre"
                                            className="text-lg font-medium text-muted-foreground hover:text-foreground transition-colors"
                                        >
                                            Sobre
                                        </Link>
                                        <Link
                                            href="#contact"
                                            className="text-lg font-medium text-muted-foreground hover:text-foreground transition-colors"
                                        >
                                            Contato
                                        </Link>
                                    </nav>
                                    <div className="space-y-2 border-t pt-4">
                                        <Button variant="outline" className="w-full" asChild>
                                            <Link href="/login">Log in</Link>
                                        </Button>
                                        <Button className="w-full" asChild>
                                            <Link href="/signup">Sign up</Link>
                                        </Button>
                                    </div>
                                </>
                            )}
                        </div>
                    </SheetContent>
                </Sheet>
            </div>
        </header >
    )
}

