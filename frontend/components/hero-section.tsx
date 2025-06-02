import { Button } from "@/components/ui/button"
import { CheckSquare, Zap } from "lucide-react"
import Link from "next/link"

export default function HeroSection() {
    return (
        <section className="relative overflow-hidden bg-gradient-to-br from-blue-50 via-white to-indigo-50 py-20 sm:py-32">
            <div className="container mx-auto px-4">
                <div className="mx-auto max-w-4xl text-center">
                    {/* Badge */}
                    <div className="inline-flex items-center rounded-full bg-blue-100 px-4 py-2 text-sm font-medium text-blue-700 mb-8">
                        <Zap className="mr-2 h-4 w-4" />
                        Correção automática em segundos
                    </div>

                    {/* Heading */}
                    <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl lg:text-7xl">
                        Corrija gabaritos
                        <span className="text-blue-600"> automaticamente</span>
                    </h1>

                    {/* Description */}
                    <p className="mt-6 text-lg leading-8 text-gray-600 sm:text-xl max-w-3xl mx-auto">
                        Transforme a correção de provas em um processo rápido e preciso. O GradeFlow usa inteligência artificial
                        para corrigir gabaritos instantaneamente, economizando horas de trabalho manual.
                    </p>

                    {/* CTA Buttons */}
                    <div className="mt-10 flex flex-col sm:flex-row items-center justify-center gap-4">
                        <Button size="lg" className="text-lg px-8 py-4" asChild>
                            <Link href="/cadastro">
                                Começar Grátis
                                <CheckSquare className="ml-2 h-5 w-5" />
                            </Link>
                        </Button>
                        <Button variant="outline" size="lg" className="text-lg px-8 py-4" asChild>
                            <Link href="#demo">Ver Demonstração</Link>
                        </Button>
                    </div>

                    {/* Stats */}
                    <div className="mt-16 grid grid-cols-1 gap-8 sm:grid-cols-3 lg:gap-16">
                        <div className="text-center">
                            <div className="text-3xl font-bold text-blue-600">99.9%</div>
                            <div className="text-sm text-gray-600">Precisão na correção</div>
                        </div>
                        <div className="text-center">
                            <div className="text-3xl font-bold text-blue-600">5s</div>
                            <div className="text-sm text-gray-600">Tempo médio de correção</div>
                        </div>
                        <div className="text-center">
                            <div className="text-3xl font-bold text-blue-600">10k+</div>
                            <div className="text-sm text-gray-600">Professores ativos</div>
                        </div>
                    </div>
                </div>
            </div>

            {/* Background decoration */}
            <div className="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80">
                <div className="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-blue-400 to-indigo-600 opacity-20 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"></div>
            </div>
        </section>
    )
}

