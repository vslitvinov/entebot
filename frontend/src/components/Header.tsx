import React from 'react';

export function Header() {
    return (
        <div className="relative bg-white">
            <div className="mx-auto max-w-7xl px-4 sm:px-6">
                <div className="flex items-center justify-between py-6 md:justify-start md:space-x-10">

                {/* start logo */}
                    <div className="flex justify-start lg:w-0 lg:flex-1">
                        <a href="#">
                            <span className="sr-only">Eterinte</span>
                            <img className="h-16 w-auto sm:h-12" src="./images/logo.png" alt="" />
                        </a>
                    </div>
                {/* end logo */}

                {/* start list items */}
                    <nav className=" space-x-10 md:flex">

                        <div className="relative">
                             <button type="button" className="group inline-flex items-center rounded-md p-1 text-base font-medium hover:bg-indigo-200 focus:text-white focus:bg-indigo-500" aria-expanded="false">
                                <span>Trading Bots</span>
                            </button>
                            <div className="absolute z-10 -ml-4 mt-3 w-auto max-w-md transform px-2 sm:px-0 lg:left-1/2 lg:ml-0 lg:-translate-x-1/2">
                                <div className="overflow-hidden rounded-lg shadow-lg ring-1 ring-black ring-opacity-5">
                                    <div className="relative grid gap-6 bg-white px-2 py-6 sm:gap-8 sm:p-8">
                                        <HeaderListItem />
                                        <HeaderListItem />
                                        <HeaderListItem />
                                    </div> 
                                </div> 
                            </div> 
                        </div> 

                    </nav>
                    <div className="hidden items-center justify-end md:flex md:flex-1 lg:w-0">
                        <a href="#" className="whitespace-nowrap text-base font-medium text-gray-500 hover:text-gray-900">Sign in</a>
                        <a href="#" className="ml-8 inline-flex items-center justify-center whitespace-nowrap rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-indigo-700">Try It Free</a>
                    </div>
                {/* end list items */}




                </div>        
            </div>        
        </div>        
    )
};

export function HeaderListItem() {
    return (
        <a href="#" className="-m-3 flex items-start rounded-lg p-2 hover:bg-gray-50">
            <svg className="h-6 w-6 flex-shrink-0 text-indigo-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z" />
            </svg>
            <div className="ml-4">
                <p className="text-base whitespace-nowrap font-medium text-gray-900">DCA Bot</p>
                <p className="">Purchase at opputune moments</p>
            </div>
        </a>
    )
}