#pragma once

#include "Game.h"
#include "Maze.h"
#include "MazeSection.h"
#include <string>
#include <memory>

namespace my {

	//Cria uma nova inst�ncia do labirinto padr�o do jogo
	MazePtr buildDefaultMaze(my::GamePtr g);

	//Cria um novo conjunto de se��es para um labirinto com base em um mapa de caracteres
	MazeSectionMatrix buildMazeSections(const std::string& referenceMap, int& width, int& height );

	//Cria um novo conjunto de se��es para um labirinto com largura e altura informados, onde todas as se��es est�o interligadas e s�o consideradas acess�veis por padr�o
	MazeSectionMatrix buildMazeSections(int width, int height);
}