import React from 'react';

export function Tools() {
    return (
    <main className="">
        <div className="relative">
            <div className="mx-auto max-w-3xl pt-4 pb-32 sm:pt-48 sm:pb-40">
                <div>
                    <div>
                        <h6 className="text-black  text-4xl font-bold tracking-tight sm:text-center sm:text-6xl">Traiding Tools</h6>
                        <p className="mt-6 text-base leading-8 text-black sm:text-center">Choose the best strategy to meet your plans and expectations</p>
                    </div>
                </div>
            </div>
            <div className="grid">
                <ToolItem />
                <ToolItem />
                <ToolItem />
                <ToolItem />
                <ToolItem />
                <ToolItem />
            </div>
        </div>
    </main>

    );
}


export function ToolItem() {
    return (
        <div className="relative">

            <div>
                <img src="" alt="" />
                <h6>DCA Bot</h6>
                <p>No need to risk it all! Instead of investing a lump sum with unknown risks, the bot will invest the amount partially with maximum benefit.</p>
                <a href="#">
                Read More 
                    <svg className="text-eblack ml-2 h-5 w-5 group-hover:text-gray-500 -rotate-90" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                        <path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
                    </svg>
                </a>
            </div>

        </div>
    );
}

