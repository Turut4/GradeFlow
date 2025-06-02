import HeroSection from '@/components/hero-section';
import FeaturesSection from '@/components/features-section';
import HowItWorks from '@/components/how-it-works';
import PricingSection from '@/components/pricing-section';
import Footer from '@/components/footer';
import PublicHeader from '@/components/public-header';

export default function LandingPage() {
    return (
        <div className="min-h-screen bg-white">
            <PublicHeader currentPage='/landing-page' />
            <main>
                <HeroSection />
                <FeaturesSection />
                <HowItWorks />
                <PricingSection />
            </main>
            <Footer />
        </div>
    );
}
