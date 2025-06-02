import Link from "next/link"
import { CheckSquare, Mail, Phone, MapPin } from "lucide-react"

export default function Footer() {
    return (
        <footer className="bg-gray-900 text-white">
            <div className="container mx-auto px-4 py-16">
                <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
                    {/* Logo and Description */}
                    <div className="col-span-1 md:col-span-2">
                        <div className="flex items-center space-x-2 mb-4">
                            <div className="flex items-center justify-center w-8 h-8 bg-blue-600 rounded-lg">
                                <CheckSquare className="w-5 h-5 text-white" />
                            </div>
                            <span className="text-xl font-bold">GradeFlow</span>
                        </div>
                        <p className="text-gray-400 mb-6 max-w-md">
                            Revolucionando a educação com correção automática de gabaritos. Mais tempo para ensinar, menos tempo
                            corrigindo.
                        </p>
                        <div className="space-y-2">
                            <div className="flex items-center text-gray-400">
                                <Mail className="w-4 h-4 mr-2" />
                                contato@gradeflow.com.br
                            </div>
                            <div className="flex items-center text-gray-400">
                                <Phone className="w-4 h-4 mr-2" />
                                (11) 9999-9999
                            </div>
                            <div className="flex items-center text-gray-400">
                                <MapPin className="w-4 h-4 mr-2" />
                                São Paulo, SP
                            </div>
                        </div>
                    </div>

                    {/* Product Links */}
                    <div>
                        <h3 className="text-lg font-semibold mb-4">Produto</h3>
                        <ul className="space-y-2">
                            <li>
                                <Link href="#recursos" className="text-gray-400 hover:text-white transition-colors">
                                    Recursos
                                </Link>
                            </li>
                            <li>
                                <Link href="#precos" className="text-gray-400 hover:text-white transition-colors">
                                    Preços
                                </Link>
                            </li>
                            <li>
                                <Link href="/demo" className="text-gray-400 hover:text-white transition-colors">
                                    Demonstração
                                </Link>
                            </li>
                            <li>
                                <Link href="/api" className="text-gray-400 hover:text-white transition-colors">
                                    API
                                </Link>
                            </li>
                        </ul>
                    </div>

                    {/* Support Links */}
                    <div>
                        <h3 className="text-lg font-semibold mb-4">Suporte</h3>
                        <ul className="space-y-2">
                            <li>
                                <Link href="/ajuda" className="text-gray-400 hover:text-white transition-colors">
                                    Central de Ajuda
                                </Link>
                            </li>
                            <li>
                                <Link href="/contato" className="text-gray-400 hover:text-white transition-colors">
                                    Contato
                                </Link>
                            </li>
                            <li>
                                <Link href="/status" className="text-gray-400 hover:text-white transition-colors">
                                    Status do Sistema
                                </Link>
                            </li>
                            <li>
                                <Link href="/blog" className="text-gray-400 hover:text-white transition-colors">
                                    Blog
                                </Link>
                            </li>
                        </ul>
                    </div>
                </div>

                <div className="border-t border-gray-800 mt-12 pt-8 flex flex-col md:flex-row justify-between items-center">
                    <p className="text-gray-400 text-sm">© 2024 GradeFlow. Todos os direitos reservados.</p>
                    <div className="flex space-x-6 mt-4 md:mt-0">
                        <Link href="/privacidade" className="text-gray-400 hover:text-white text-sm transition-colors">
                            Política de Privacidade
                        </Link>
                        <Link href="/termos" className="text-gray-400 hover:text-white text-sm transition-colors">
                            Termos de Uso
                        </Link>
                    </div>
                </div>
            </div>
        </footer>
    )
}

