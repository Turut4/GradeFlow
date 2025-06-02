import AuthForm from '@/components/auth-form';
import PublicHeader from '@/components/public-header'

export default function RegisterPage() {
    return (

        <div className="min-h-screen bg-white">
            <PublicHeader currentPage='/register' />
            <main className="flex justify-center items-start pt-20 min-h-screen bg-gray-50">
                <AuthForm type="register" />
            </main>
        </div>
    );
}
