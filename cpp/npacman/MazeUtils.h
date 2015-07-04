#pragma once

#include "Game.h"
#include "Maze.h"
#include "MazeSection.h"
#include <string>
#include <memory>

namespace my {

	//Cria uma nova instância do labirinto padrão do jogo
	MazePtr buildDefaultMaze(my::GamePtr g);

	//Cria um novo conjunto de seções para um labirinto com base em um mapa de caracteres
	MazeSectionMatrix buildMazeSections(const std::string& referenceMap, int& width, int& height );

	//Cria um novo conjunto de seções para um labirinto com largura e altura informados, onde todas as seções estão interligadas e são consideradas acessíveis por padrão
	MazeSectionMatrix buildMazeSections(int width, int height);
}