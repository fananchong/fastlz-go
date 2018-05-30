workspace "ctest"
    configurations { "Debug", "Release" }
    platforms { "x64" }
    targetdir "../bin/"
    language "C++"
    includedirs {
        "..",
        "../Contrib/fastlz",
    }
    flags {
        "C++11",
        "StaticRuntime",
    }

    filter "configurations:Debug"
    defines { "_DEBUG" }
    flags { "Symbols" }
    libdirs { }
    filter "configurations:Release"
    defines { "NDEBUG" }
    libdirs { }
    optimize "On"
    filter { }  


project "ctest"
    kind "ConsoleApp"
    targetname "ctest"
    libdirs { "../bin" }
    files {
        "../main.cpp",
        "../Contrib/fastlz/**",
    }
--[[    
project "cbenchmark"
    kind "ConsoleApp"
    targetname "cbenchmark"
    libdirs { "../bin" }
    files {
        "../benchmark.cpp",
        "../Contrib/fastlz/**",
    }
]]--