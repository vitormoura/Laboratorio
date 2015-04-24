#pragma once
#include <string>
#include <memory>

namespace my {

	class ResourceManager
	{

	public:
		ResourceManager();
		~ResourceManager();

		//Recupera conte�do do mapa de caracteres que auxilia na constru��o de labirintos
		std::unique_ptr<std::string> getDefaultMazeBlueprints();

		//Recupera conte�do de um arquivo de texto na forma de uma string
		std::unique_ptr<std::string> getFileContents(const char* filePath);
	};
}

