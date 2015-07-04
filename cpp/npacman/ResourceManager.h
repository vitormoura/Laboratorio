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

		//Recupera arquivo de textura padrão par ao background o labirinto
		sf::Texture* ResourceManager::getDefaultMazeTemplate();

		//Recupera conteúdo do mapa de caracteres que auxilia na construção de labirintos
		std::unique_ptr<std::string> getDefaultMazeBlueprints();

		//Recupera conteúdo de um arquivo de texto na forma de uma string
		std::unique_ptr<std::string> getFileContents(const char* filePath);
	};
}

