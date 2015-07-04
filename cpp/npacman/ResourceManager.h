#pragma once
#include <string>
#include <memory>
#include <SFML\Graphics.hpp>

namespace my {

	class ResourceManager
	{
	private:
		sf::Texture* m_default_maze_bg;

	public:
		ResourceManager();
		~ResourceManager();

		//Recupera arquivo de textura padr�o par ao background o labirinto
		sf::Texture* ResourceManager::getDefaultMazeTemplate();

		//Recupera conte�do do mapa de caracteres que auxilia na constru��o de labirintos
		std::unique_ptr<std::string> getDefaultMazeBlueprints();

		//Recupera conte�do de um arquivo de texto na forma de uma string
		std::unique_ptr<std::string> getFileContents(const char* filePath);
	};
}

