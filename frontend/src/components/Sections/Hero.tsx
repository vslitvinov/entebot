import React from 'react';

export function Hero() {
    return (
        <main className="bg-hero ">
        <div className="bg-herobg/[.45] relative px-6 p-16 lg:px-8">
            <div className="mx-auto max-w-3xl pt-4 pb-32 sm:pt-48 sm:pb-40">
                <div>
                    <div>
                        <h1 className="text-white  text-4xl font-bold tracking-tight sm:text-center sm:text-6xl">Automate your trading</h1>
                        <p className="mt-6 text-lg leading-8 text-white sm:text-center">The best way to make a profit regardless of the market situation</p>
                        <div className="mt-8 flex gap-x-4 sm:justify-center">
                            <a href="#" className="inline-block rounded-lg bg-indigo-600 px-4 py-1.5 text-base font-semibold leading-7 text-white shadow-sm ring-1 ring-indigo-600 hover:bg-indigo-700 hover:ring-indigo-700">
                                Get started
                                <span className="text-indigo-200" aria-hidden="true">&rarr;</span>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>

    );
}

