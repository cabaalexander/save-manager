import { IconButton } from "@material-tailwind/react";
import clsx from "clsx";
import { FaPlus } from "react-icons/fa6";
import Game from "@/components/Game";
import useMenuMiddleItem from "@/hooks/useMenuMiddleItem";

const Saves = () => {
  useMenuMiddleItem(
    <IconButton data-testid="addGameSave">
      <FaPlus />
    </IconButton>,
  );

  return (
    <main>
      <ul className={clsx("mb-4", "flex flex-wrap justify-around gap-3")}>
        <Game title="DS1" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
        <Game title="Game" />
      </ul>
    </main>
  );
};

export default Saves;
