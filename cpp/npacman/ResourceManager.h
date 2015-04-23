#pragma once
#include <string>
#include <memory>

namespace my {

	class ResourceManager
	{

	public:
		ResourceManager();
		~ResourceManager();

		//Recupera conteúdo de um arquivo de texto na forma de uma string
		std::unique_ptr<std::string> getFileContents(const char* filePath);
	};
}

