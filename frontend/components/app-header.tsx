import { usePathname } from "next/navigation"
import { Button } from "@/components/ui/button"
import Link from "next/link"
import { useAuthActions } from "@/hooks/use-auth-actions"
import { LogOut, Settings, User } from "lucide-react"
import { navigationItems } from "@/lib/constants/navigation"

export default function AppHeader() {
    const { logout } = useAuthActions()
    const pathname = usePathname()

    const handleLogout = () => logout()

    return (
        <header className="sticky top-0 z-50 flex h-16 items-center justify-between bg-background px-4 shadow-sm">
            <div className="flex items-center space-x-4">
                <Link href="/dashboard" className="text-xl font-bold">GradeFlow</Link>

                <nav className="hidden md:flex space-x-2">
                    {navigationItems.map((item) => (
                        <Link
                            key={item.href}
                            href={item.href}
                            className={`text-sm px-3 py-2 rounded-md ${pathname === item.href
                                ? "bg-accent text-accent-foreground"
                                : "text-muted-foreground hover:bg-accent hover:text-accent-foreground"
                                }`}
                        >
                            {item.label}
                        </Link>
                    ))}
                </nav>
            </div>

            <div className="flex items-center space-x-2">
                <Link href="/settings">
                    <Button variant="ghost" size="icon"><Settings className="w-4 h-4" /></Button>
                </Link>
                <Link href="/profile">
                    <Button variant="ghost" size="icon"><User className="w-4 h-4" /></Button>
                </Link>
                <Button onClick={handleLogout} variant="outline" size="icon">
                    <LogOut className="w-4 h-4" />
                </Button>
            </div>
        </header>
    )
}

