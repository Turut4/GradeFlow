'use client';
import AuthForm from '@/components/auth-form';
import PublicHeader from '@/components/public-header';

export default function LoginPage() {
    return (
        <div className="min-h-screen bg-white">
            <PublicHeader currentPage='/login' />
            <main className="flex justify-center items-start pt-40 h-screen bg-gray-50 ">
                <AuthForm type="login" />
            </main>
        </div>
    );
}
